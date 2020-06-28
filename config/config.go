package config

import "github.com/spf13/viper"

var config *DefaultConfig

type DefaultConfig struct {
	LocalViper  *viper.Viper
	RemoteViper *viper.Viper
	*viper.Viper
}

func GetString(key string) string {
	return config.GetString(key)
}

func (config *DefaultConfig) GetString(key string) string {
	if value := config.LocalViper.GetString(key); value != "" {
		return value
	}
	if value := config.RemoteViper.GetString(key); value != "" {
		return value
	}
	return config.Viper.GetString(key)
}
