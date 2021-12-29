package config

import "github.com/spf13/viper"

func Init() error {
	viper.AddConfigPath("./config")
	viper.SetConfigType("json")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
