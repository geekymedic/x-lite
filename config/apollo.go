package config

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/shima-park/agollo"
)

type apolloBackend struct {
	ctx        context.Context
	cancel     context.CancelFunc
	once       *sync.Once
	watchEvent chan Event
	clis       map[string]agollo.Agollo
	cache      map[string][]byte
}

func newApolloBackend(appid, endpoint, path string) (*apolloBackend, error) {
	ctx, cancel := context.WithCancel(context.Background())
	var clis = make(map[string]agollo.Agollo)
	for _, nameSpace := range strings.Split(path, "@") {
		cli, err := agollo.New(endpoint, appid, agollo.AutoFetchOnCacheMiss(), agollo.FailTolerantOnBackupExists())
		if err != nil {
			return nil, err
		}
		cli.Start()
		clis[nameSpace] = cli
		fmt.Fprintln(os.Stdout, "Appo namespace")
	}

	return &apolloBackend{clis: clis, ctx: ctx, cancel: cancel, once: &sync.Once{}, watchEvent: make(chan Event)}, nil
}

func (backend *apolloBackend) Get(_ context.Context, _ string) ([]byte, error) {
	var rd = bytes.NewBuffer(nil)
	for nameSpace, cli := range backend.clis {
		conf := cli.GetNameSpace(nameSpace)
		if content, ok := conf["content"]; ok {
			rd.WriteString(content.(string))
			rd.WriteString("\n")
		}
	}
	return rd.Bytes(), nil
}

func (backend *apolloBackend) Watch(_ context.Context, _ string) <-chan Event {
	backend.once.Do(func() {
		for _, cli := range backend.clis {
			go func(cli agollo.Agollo) {
				watchChan := cli.Watch()
				for {
					select {
					case <-backend.ctx.Done():
						return
					case ev := <-watchChan:
						if _, ok := backend.clis[ev.Namespace]; ok {
							backend.watchEvent <- Event{
								EV: ev,
							}
						}
					}
				}
			}(cli)
			break
		}
	})
	return backend.watchEvent
}

func (backend *apolloBackend) Close() error {
	backend.cancel()
	return nil
}
