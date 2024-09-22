package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	LogLevel      string `yaml:"logLevel"`
	ServerAddress string `yaml:"serverAddress"`
}

func Load() (*Config, error) {
	viper.SetConfigFile("config.yaml")
	viper.SetEnvPrefix("MY_APP")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, err
		}
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
