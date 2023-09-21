package config

import "github.com/spf13/viper"

type DBConfig struct {
	User    string
	Pass    string
	Host    string
	Port    int
	Name    string
	SSLMODE string
}

func NewDBConfig(v *viper.Viper) *DBConfig {
	return &DBConfig{
		User:    v.GetString("DB_USER"),
		Pass:    v.GetString("DB_PASS"),
		Host:    v.GetString("DB_HOST"),
		Port:    v.GetInt("DB_PORT"),
		Name:    v.GetString("DB_NAME"),
		SSLMODE: v.GetString("DB_SSLMODE"),
	}
}
