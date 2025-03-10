package config

import "github.com/spf13/viper"

type LogConfig interface {
	Level() int
}

type logCfg struct {
	viper *viper.Viper
}

func (l *logCfg) Level() int {
	return l.viper.GetInt("Log_Level")
}

func newLogConfg(v *viper.Viper) *logCfg {
	v.SetDefault("Log_Level", 1)

	return &logCfg{v}
}
