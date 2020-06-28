package neon

import (
	"fmt"

	"github.com/spf13/viper"

	"github.com/geekymedic/neon/errors"
	"github.com/geekymedic/neon/logger"
)

type PluginStatus int

const (
	PluginLoad PluginStatus = 1
)

type PluginHandler func(status PluginStatus, viper *viper.Viper) error

var (
	_plugins = make(map[string]PluginHandler, 0)
)

func AddPlugin(name string, handler PluginHandler) {

	_, ok := _plugins[name]

	if ok {
		panic(
			fmt.Sprintf("plugin %s already exists.", name))
	}

	_plugins[name] = handler
}

func LoadPlugins(viper *viper.Viper) error {

	for name, plugin := range _plugins {

		err := plugin(PluginLoad, viper)

		if err != nil {
			return errors.WithMessage(err, "load plugin %s fail", name)
		}

		logger.With("plugin", name).Info("plugin loaded")
	}

	return nil
}
