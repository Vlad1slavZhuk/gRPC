package config

import "os"

type Config struct {
	Port string
}

func NewConfigFromEnv() *Config {
	var ok bool
	config := Config{}
	if config.Port, ok = os.LookupEnv("SERVER_PORT"); !ok {
		config.Port = "80"
	}
	config.Port = ":" + config.Port

	return &config
}
