package config

import "os"

func Load() (*Config, error) {
	cfg := &Config{}

	cfg.LogLevel = os.Getenv("LOG_LEVEL")
	if cfg.LogLevel == "" {
		cfg.LogLevel = "info"
	}

	cfg.Database.DSN = os.Getenv("DATABASE_DSN")
	if cfg.Database.DSN == "" {
		cfg.Database.DSN = "postgres://postgres:postgres@localhost:5432/syncra?sslmode=disable"
	}

	cfg.Database.MaxOpenConns = 20
	cfg.Database.MaxIdleConns = 10

	return cfg, nil
}

type Config struct {
	LogLevel string
	Database struct {
		DSN          string
		MaxOpenConns int
		MaxIdleConns int
	}
}
