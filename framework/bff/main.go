package bff

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/viper"

	"git.gmtshenzhen.com/geeky-medic/x-lite/config"
	"git.gmtshenzhen.com/geeky-medic/x-lite/logger"
	"git.gmtshenzhen.com/geeky-medic/x-lite/version"
	errors "git.gmtshenzhen.com/geeky-medic/x-lite/xerrors"
)

func Main() error {
	fmt.Fprintln(os.Stdout, version.Version())
	conf := flag.String("c", "config", "config path")
	flag.Parse()
	err := config.LoadRemote(*conf)
	if err != nil {
		return errors.By(err)
	}
	var addr = viper.GetString("Address")
	srv := &http.Server{
		Addr:    addr,
		Handler: nil,
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		logger.With("address", srv.Addr).Info("start listen...")
		err = srv.ListenAndServe()
		if err != nil {
			logger.Errorf("listen %s fail, %v\n", addr, err)
			err = errors.By(err)
		} else {
			logger.Infof("listen %s", srv.Addr)
		}
		c <- syscall.SIGTERM
	}()
	<-c

	logger.Info("shutdown server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return errors.By(srv.Shutdown(ctx))
}
