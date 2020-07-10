package xdb

import (
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"

	db "github.com/geekymedic/x-lite/pkg/xmysql"
	"github.com/geekymedic/x-lite/plugin"
	errors "github.com/geekymedic/x-lite/xerrors"
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

	plugin.AddPlugin("db", func(status plugin.Status, viper *viper.Viper) error {
		switch status {
		case plugin.Load:

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
