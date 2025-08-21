package config

import (
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

type (
	Config struct {
		GRPC        `envPrefix:"GRPC_"`
		Gateway     `envPrefix:"GATEWAY_"`
		TaskTracker `envPrefix:"TASKTRACKER_"`
	}

	GRPC struct {
		Host string `env:"HOST"`
		Port string `env:"PORT"`
	}

	Gateway struct {
		Host string `env:"HOST"`
		Port string `env:"PORT"`
	}

	TaskTracker struct {
		Host string `env:"HOST"`
		Port string `env:"PORT"`
	}
)

func New() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Print("warning: no .env file")
	}

	cfg := &Config{}

	if err := env.Parse(cfg); err != nil {
		return nil, errors.Wrap(err, "parse .env")
	}

	return cfg, nil
}
