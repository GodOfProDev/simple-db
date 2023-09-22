package config

import "github.com/spf13/viper"

type DBConfig struct {
	Url string
}

func NewDBConfig(v *viper.Viper) *DBConfig {
	return &DBConfig{
		Url: v.GetString("DB_URL"),
	}
}
