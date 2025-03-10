package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Cfg struct {
	Server ServerConfig
	Log    LogConfig
}

func NewConfig() *Cfg {
	v := loadViper()

	c := &Cfg{
		Server: newServerConfig(v),
		Log:    newLogConfg(v),
	}

	return c
}

func loadViper() *viper.Viper {
	v := viper.New()

	v.AddConfigPath(".")
	v.SetEnvPrefix("puck")
	v.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		} else {
			panic(fmt.Errorf("fatal error config file: %w", err))
		}
	}

	return v
}
