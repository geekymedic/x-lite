package xredis

import (
	"strings"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"

	"github.com/geekymedic/x-lite/plugin"
	errors "github.com/geekymedic/x-lite/xerrors"
)

var (
	redisList                = map[string]*redis.Client{}
	handler   plugin.Handler = func(status plugin.Status, viper *viper.Viper) error {
		type RedisOptions struct {
			DSN      string
			Password string
			DB       int
		}
		switch status {
		case plugin.Load:
			var (
				dsnList = make(map[string]*RedisOptions)
			)

			err := viper.UnmarshalKey("redis", &dsnList)

			if err != nil {
				return errors.By(err)
			}

			if len(dsnList) == 0 {
				return errors.Format("redis plugin used, but redis config not exists.")
			}

			for name, opt := range dsnList {

				client := redis.NewClient(&redis.Options{
					Addr:     opt.DSN,
					Password: opt.Password,
					DB:       opt.DB,
				})

				redisList[strings.ToLower(name)] = client
			}

		}
		return nil
	}
)

func init() {
	plugin.AddPlugin("redis", handler)
}

func Use(name string) *redis.Client {
	return redisList[strings.ToLower(name)]
}
