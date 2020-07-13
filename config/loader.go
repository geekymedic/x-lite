package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"

	"github.com/geekymedic/x-lite/logger"
	errors "github.com/geekymedic/x-lite/xerrors"
)

const (
	XLiteModeLt      = "LT"
	XLiteModeDev     = "DEV"
	XLiteModeTest    = "TEST"
	XLiteModeProduct = "PRD"
)

const (
	XLiteMode           = "X-LITE_MODE"
	XLiteConfigProvider = "X-LITE_CONFIG_PROVIDER"
	XLiteConfigEndpoint = "X-LITE_CONFIG_ENDPOINT"
	XLiteConfigPath     = "X-LITE_CONFIG_PATH"
	XLiteConfigSecret   = "X-LITE_CONFIG_SECRET"
)

func init() {
	viper.BindEnv(XLiteMode)
	viper.BindEnv(XLiteConfigProvider)
	viper.BindEnv(XLiteConfigEndpoint)
	viper.BindEnv(XLiteConfigPath)
	viper.BindEnv(XLiteConfigSecret)
}

var RemoterViper = viper.New()
var LocalViper = viper.New()

// Must compatible old version, don't' ignore the function
func Load(path *string) error {
	if env := viper.GetString(XLiteMode); env != "" {
		viper.SupportedRemoteProviders = []string{"etcd", "apollo"}
		env = strings.ToUpper(env)
		if env != XLiteModeLt && env != XLiteModeDev && env != XLiteModeTest && env != XLiteModeProduct {
			return errors.Format("x-lite env should be set %s OR %s OR %s OR %s", XLiteModeLt, XLiteModeDev, XLiteModeTest, XLiteModeProduct)
		}
		if err := LoadRemote(*path); err != nil {
			return errors.By(err)
		}

		return nil
	}

	viper.AddConfigPath(*path)
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	err := viper.ReadInConfig()
	if err != nil {
		return errors.By(err)
	}
	LocalViper = viper.GetViper()
	viper.WatchConfig()
	return nil
}

func LoadRemote(localPath string) error {
	// read the local config
	{
		LocalViper.AddConfigPath(localPath)
		LocalViper.SetConfigName("config")
		LocalViper.SetConfigType("yml")
		if err := LocalViper.ReadInConfig(); err != nil {
			return errors.By(err)
		}
		if err := LocalViper.ReadInConfig(); err != nil {
			return errors.By(err)
		}
		LocalViper.WatchConfig()
	}
	viper.SetConfigType("yml")
	logger.With(XLiteConfigProvider, viper.Get(XLiteConfigProvider),
		XLiteConfigEndpoint, viper.Get(XLiteConfigEndpoint),
		XLiteConfigPath, viper.Get(XLiteConfigPath),
		XLiteConfigSecret, viper.Get(XLiteConfigSecret)).Debug("config env trace")
	err := viper.AddSecureRemoteProvider(viper.GetString(XLiteConfigProvider),
		viper.GetString(XLiteConfigEndpoint),
		viper.GetString(XLiteConfigPath),
		viper.GetString(XLiteConfigSecret))
	if err != nil {
		return errors.By(err)
	}
	if err := viper.MergeConfigMap(LocalViper.AllSettings()); err != nil {
		return errors.By(err)
	}
	if err = viper.ReadRemoteConfig(); err != nil {
		return errors.By(err)
	}
	if err = viper.MergeConfigMap(LocalViper.AllSettings()); err != nil {
		return errors.By(err)
	}
	if err = viper.GetViper().WatchRemoteConfigOnChannel(); err != nil {
		return errors.By(err)
	}
	// Bind remote viper
	RemoterViper = viper.GetViper()
	fmt.Fprintf(os.Stdout, "Load remote config, provider is %s\n", XLiteConfigProvider)
	return nil
}
