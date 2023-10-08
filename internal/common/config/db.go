package config

import (
	"github.com/caarlos0/env/v9"
)

type DBConfig struct {
	DBHost     string `env:"GIZLOG_DB_HOST"`
	DBPort     int    `env:"GIZLOG_DB_PORT"`
	DBUser     string `env:"GIZLOG_DB_USER"`
	DBPassword string `env:"GIZLOG_DB_PASSWORD"`
	DBName     string `env:"GIZLOG_DB_NAME"`
}

func NewDBConfig() (*DBConfig, error) {
	cfg := &DBConfig{}
	if err := env.Parse(cfg); err != nil {
		return &DBConfig{}, err
	}

	return cfg, nil
}
