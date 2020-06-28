package config

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"strings"
	"sync"
	"time"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"

	errors "git.gmtshenzhen.com/geeky-medic/x-lite/xerrors"
)

type Event struct {
	EV interface{}
}

type Backend interface {
	Get(ctx context.Context, path string) ([]byte, error)
	Watch(ctx context.Context, path string) <-chan Event
	Close() error
}

// Only support etcd v3.0
type remoteConfigProvider struct {
	onlyOnce *sync.Once
	etcd     Backend
	apollo   Backend
}

func (rc *remoteConfigProvider) initBackend(rp viper.RemoteProvider) {
	if rp.Provider() == "etcd" {
		rc.onlyOnce.Do(func() {
			var (
				endpoint           = rp.Endpoint()
				secretKeyString    = strings.SplitN(rp.SecretKeyring(), "@", 2)
				userName, password string
				err                error
			)
			if len(secretKeyString) == 2 {
				userName = secretKeyString[0]
				password = secretKeyString[1]
			}
			var endpoints []string
			for _, endpoint := range strings.Split(endpoint, ",") {
				endpoints = append(endpoints, endpoint)
			}
			rc.etcd, err = newEtcdBackend(userName, password, endpoints, time.Second*3)
			if err != nil {
				panic(err)
			}
			fmt.Println("Init etcd provider successfully")
		})
	}
	if rp.Provider() == "apollo" {
		panic("Unsupported the provider")
		rc.onlyOnce.Do(func() {
			var err error
			rc.apollo, err = newApolloBackend(rp.SecretKeyring(), rp.Endpoint(), rp.Path())
			if err != nil {
				panic(err)
			}
			fmt.Println("Init apollo provider successfully")
		})
	}
}

// patch for local config
func (rc *remoteConfigProvider) Get(rp viper.RemoteProvider) (io.Reader, error) {
	rc.initBackend(rp)
	buf, err := rc.backend(rp.Provider()).Get(context.Background(), rp.Path())
	if err != nil {
		return nil, errors.By(err)
	}
	if keyValues := LocalViper.AllSettings(); len(keyValues) > 0 {
		patch, err := yaml.Marshal(LocalViper.AllSettings())
		if err != nil {
			return nil, errors.By(err)
		}
		buf = append(buf, patch...)
	}

	return bytes.NewReader(buf), nil
}

func (rc *remoteConfigProvider) Watch(rp viper.RemoteProvider) (io.Reader, error) {
	rc.initBackend(rp)
	buf, err := rc.backend(rp.Provider()).Get(context.Background(), rp.Path())
	if err != nil {
		return nil, errors.By(err)
	}

	return bytes.NewReader(buf), nil
}

func (rc *remoteConfigProvider) WatchChannel(rp viper.RemoteProvider) (<-chan *viper.RemoteResponse, chan bool) {
	rc.initBackend(rp)
	quitwc := make(chan bool)
	viperResponseCh := make(chan *viper.RemoteResponse)

	var backend = rc.backend(rp.Provider())

	go func(vr <-chan *viper.RemoteResponse, quitwc <-chan bool) {
		defer backend.Close()
		for {
			select {
			case <-quitwc:
				return
			case <-backend.Watch(context.Background(), rp.Path()):
				fmt.Println("Receive a config change event ", rp.Path())
				buf, err := backend.Get(context.Background(), rp.Path())
				viperResponseCh <- &viper.RemoteResponse{
					Error: err,
					Value: buf,
				}
			}
		}
	}(viperResponseCh, quitwc)
	return viperResponseCh, quitwc
}

func (rc *remoteConfigProvider) backend(provider string) Backend {
	if provider == "etcd" {
		return rc.etcd
	}
	if provider == "apollo" {
		return rc.apollo
	}
	return nil
}

func init() {
	viper.RemoteConfig = &remoteConfigProvider{
		onlyOnce: &sync.Once{},
	}
}
