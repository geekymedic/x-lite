package xmetrics

import (
	_ "net/http/pprof"

	"github.com/spf13/viper"

	"github.com/geekymedic/x-lite/logger"
	"github.com/geekymedic/x-lite/metrics/prometheus"
	"github.com/geekymedic/x-lite/plugin"
	errors "github.com/geekymedic/x-lite/xerrors"
)

func init() {
	plugin.AddPlugin("metrics", func(status plugin.Status, viper *viper.Viper) error {
		switch status {
		case plugin.Load:

			addr := viper.GetString("Metrics.Address")
			if addr == "" {
				logger.Warn("not found Metrics.Address config")
				return nil
			}
			path := viper.GetString("Metrics.Path")
			if path == "" {
				path = "/metrics"
			}

			if len(addr) == 0 {
				return errors.Format("load Metrics.Address fail, empty address.")
			}
			go func() {
				logger.With("lis", addr, "path", path).Info("start metrics server")
				err := prometheus.StartMetricsServer(addr, path)
				if err != nil {
					logger.Error("fail to start metrics server", "err", err)
				}
			}()
		}

		return nil
	})

}
