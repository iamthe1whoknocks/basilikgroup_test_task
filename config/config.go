package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Host string
	Port string
}

func New() *Config {
	return &Config{}
}

func (c *Config) Init() error {

	viper.AddConfigPath("../config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()

	return err
}
