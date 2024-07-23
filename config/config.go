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

func ProvideConfig() (*Config, error) {
	var cfg Config

	if err := loadEnv(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func loadEnv(config *Config) error {

	config.Server.Port = os.Getenv("PORT")

	if config.Server.Port == "" {
		return fmt.Errorf("missing environment variable PORT")
	}

	config.MongoDB.URI = os.Getenv("MONGO_URI")

	if config.MongoDB.URI == "" {
		return fmt.Errorf("missing environment variable MONGO_URI")
	}

	config.MongoDB.Database = os.Getenv("MONGO_DATABASE")

	if config.MongoDB.Database == "" {
		return fmt.Errorf("missing environment variable MONGO_DATABASE")
	}

	return nil
}
