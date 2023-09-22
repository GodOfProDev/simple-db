package config

import "github.com/spf13/viper"

type ServerConfig struct {
	Host string
	Port int
}

func NewServerConfig(v *viper.Viper) *ServerConfig {
	return &ServerConfig{Host: v.GetString("HOST"), Port: v.GetInt("PORT")}
}
