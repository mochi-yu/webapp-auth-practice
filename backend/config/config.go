package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	Port             string `env:"PORT,notEmpty" envDefault:"8080"`
	DBPort           string `env:"DB_PORT,notEmpty"`
	DBUser           string `env:"DB_USER,notEmpty"`
	DBPassword       string `env:"DB_PASSWORD,notEmpty"`
	DBName           string `env:"DB_NAME,notEmpty"`
	DBHost           string `env:"DB_HOST,notEmpty"`
	DBDataSource     string
	CorsAllowOrigins string `env:"CORS_ALLOW_ORIGINS,notEmpty"`
}

func NewConfig() (*Config, error) {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		return nil, err
	}

	cfg.DBDataSource = fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	)

	return &cfg, nil
}
