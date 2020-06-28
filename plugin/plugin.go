package plugin

import (
	"fmt"

	"github.com/spf13/viper"

	"github.com/geekymedic/x-lite/logger"
	errors "github.com/geekymedic/x-lite/xerrors"
)

type Status int

const (
	Load Status = 1
)

type Handler func(status Status, viper *viper.Viper) error

var (
	_plugins = make(map[string]Handler, 0)
)

func AddPlugin(name string, handler Handler) {
	if _, ok := _plugins[name]; ok {
		panic(fmt.Sprintf("plugin %s already exists.", name))
	}

	_plugins[name] = handler
}

func LoadPlugins(viper *viper.Viper) error {
	for name, plugin := range _plugins {
		if err := plugin(Load, viper); err != nil {
			return errors.WithMessage(err, "load plugin %s fail", name)
		}

		logger.With("plugin", name).Info("load plugin")
	}

	return nil
}
