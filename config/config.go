package config

import (
	"fmt"
	"os"
)

type Config struct {
	Server struct {
		Port string `env:"PORT"`
	}
	MongoDB struct {
		URI      string `env:"MONGO_URI"`
		Database string `env:"MONGO_DATABASE"`
	}
}

func LoadConfig() (*Config, error) {
	var config Config

	config.Server.Port = os.Getenv("PORT")

	if config.Server.Port == "" {
		return nil, fmt.Errorf("Missing environment variable PORT")
	}

	config.MongoDB.URI = os.Getenv("MONGO_URI")

	if config.MongoDB.URI == "" {
		return nil, fmt.Errorf("Missing environment variable MONGO_URI")
	}

	config.MongoDB.Database = os.Getenv("MONGO_DATABASE")

	if config.MongoDB.Database == "" {
		return nil, fmt.Errorf("Missing environment variable MONGO_DATABASE")
	}

	return &config, nil
}
