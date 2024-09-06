package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	*viper.Viper
}

func Load(filename string) (*Config, error) {
	v := viper.New()
	v.SetConfigFile(filename)
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	return &Config{v}, nil
}