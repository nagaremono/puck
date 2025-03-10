package config

import (
	"github.com/spf13/viper"
)

type ServerConfig interface {
	Port() string
}

type serverCfg struct {
	viper *viper.Viper
}

func (s *serverCfg) Port() string {
	port := s.viper.GetString("port")

	return port
}

func newServerConfig(v *viper.Viper) *serverCfg {
	v.SetDefault("port", "8000")

	return &serverCfg{v}
}
