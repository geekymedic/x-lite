package xdb

import (
	"strings"
	"time"

	"github.com/geekymedic/neon/utils/db"

	"github.com/geekymedic/neon"
	"github.com/geekymedic/neon/errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var (
	databases = map[string]*db.DB{}
)

func init() {

	type DBOptions struct {
		DSN      string
		MaxIdle  int
		MaxOpen  int
		LifeTime int
	}

	neon.AddPlugin("xdb", func(status neon.PluginStatus, viper *viper.Viper) error {
		switch status {
		case neon.PluginLoad:

			var (
				dsnList = make(map[string]*DBOptions)
			)

			err := viper.UnmarshalKey("db", &dsnList)

			if err != nil {
				return errors.By(err)
			}

			if len(dsnList) == 0 {
				return errors.Format("xdb plugin used, but xdb config not exists.")
			}

			for name, opt := range dsnList {

				db, err := db.Open("mysql", opt.DSN)

				if err != nil {
					return errors.WithMessage(err, "open rpc [%s] by [%s] fail", name, opt.DSN)
				}

				if opt.LifeTime != 0 {
					db.SetConnMaxLifetime(time.Duration(opt.LifeTime) * time.Second)
				}

				if opt.MaxIdle != 0 {
					db.SetMaxIdleConns(opt.MaxIdle)
				}

				if opt.MaxOpen != 0 {
					db.SetMaxOpenConns(opt.MaxOpen)
				}

				databases[name] = db
			}

		}

		return nil

	})
}

func Use(name string) *db.DB {
	return databases[strings.ToLower(name)]
}
