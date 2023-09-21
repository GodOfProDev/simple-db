package config

import "github.com/spf13/viper"

type ServerConfig struct {
	Port int
}

func NewServerConfig(v *viper.Viper) *ServerConfig {
	return &ServerConfig{Port: v.GetInt("PORT")}
}
