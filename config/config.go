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
	viper.SetConfigFile("config.yaml")

	err := viper.ReadInConfig()

	return err
}
