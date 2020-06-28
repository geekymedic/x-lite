package config

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/spf13/viper"
	"go.etcd.io/etcd/clientv3"
)

type etcdBackend struct {
	ctx        context.Context
	cancel     context.CancelFunc
	cli        *clientv3.Client
	once       *sync.Once
	watchEvent chan Event
}

func newEtcdBackend(userName, password string, endpoints []string, dialTimeout time.Duration) (*etcdBackend, error) {
	ctx, cancelFn := context.WithCancel(context.Background())
	cli, err := clientv3.New(clientv3.Config{
		Username:    userName,
		Password:    password,
		Endpoints:   endpoints,
		DialTimeout: dialTimeout,
	})

	if err != nil {
		return nil, err
	}
	backend := &etcdBackend{
		ctx:        ctx,
		cancel:     cancelFn,
		cli:        cli,
		watchEvent: make(chan Event),
		once:       &sync.Once{},
	}
	return backend, nil
}

func (backend *etcdBackend) Get(ctx context.Context, path string) ([]byte, error) {
	resp, err := backend.cli.Get(ctx, path, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	var buf = bytes.NewBuffer(nil)
	for _, ev := range resp.Kvs {
		key := fmt.Sprintf("%s", ev.Key)
		keySplit := strings.Split(key, "/")
		key = keySplit[len(keySplit)-1]
		idx := strings.LastIndex(key, ".")
		if idx > 0 {
			key = key[:idx]
		}
		if strings.HasSuffix(key, viper.GetString("name")) {
			buf.Write(ev.Value)
			buf.WriteString("\n")
		} else if strings.HasPrefix(key, "public") {
			buf.Write(ev.Value)
			buf.WriteString("\n")
		}
	}
	return buf.Bytes(), nil
}

func (backend *etcdBackend) Watch(_ context.Context, path string) <-chan Event {
	backend.once.Do(func() {
		go func() {
			fmt.Println("Watch Path:", path)
			ctx := context.TODO()
			var watchChan clientv3.WatchChan
			watchChan = backend.cli.Watch(ctx, path, clientv3.WithPrefix())
			for {
				select {
				case <-backend.ctx.Done():
					return
				case ev, ok := <-watchChan:
					if !ok {
						fmt.Println("Watch chanel has closed")
						watchChan = backend.cli.Watch(ctx, path, clientv3.WithPrefix())
						continue
					}
					backend.watchEvent <- Event{
						EV: ev,
					}
				}
			}
		}()
	})
	return backend.watchEvent
}

func (backend *etcdBackend) Close() error {
	backend.cancel()
	return backend.cli.Close()
}