package config

import "os"

type Config struct {
	Port string
}

func NewConfigFromEnv() *Config {
	config := new(Config)
	config.Port = os.Getenv("PORT")
	if config.Port == "" {
		config.Port = "8080"
	}
	return config
}
