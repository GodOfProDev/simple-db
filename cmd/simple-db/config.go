package main

import "github.com/spf13/viper"

func initViper() (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigFile(".env")

	return v, v.ReadInConfig()
}
