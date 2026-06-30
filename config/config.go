package config

import (
	"errors"
	"os"
	"strconv"
	"time"
)

type ServerConfig struct {
	Host    string `mapstructure:"host"`
	Port    int    `mapstructure:"port"`
	AppName string
}

type PostgresConfig struct {
	DSN          string `mapstructure:"dsn"`
	MaxOpenConns int
	MaxIdleConns int
}

type RedisConfig struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type RabbitMQConfig struct {
	URL string `mapstructure:"url"`
}

type TokenConfig struct {
	JWTSecret string
	JWTExpiry time.Duration
}

type Config struct {
	LogLevel          string
	Server            ServerConfig
	Database          PostgresConfig
	RabbitMq          RabbitMQConfig
	Redis             RedisConfig
	Token             TokenConfig
	RateLimit         int
	RateLimitInterval time.Duration
}

func Load() (*Config, error) {
	cfg := &Config{}

	cfg.LogLevel = os.Getenv("LOG_LEVEL")
	if cfg.LogLevel == "" {
		cfg.LogLevel = "info"
	}

	cfg.Database.DSN = os.Getenv("DATABASE_DSN")
	cfg.Server.AppName = os.Getenv("APPS_NAME")
	port, err := strconv.Atoi(os.Getenv("PORT"))

	if err != nil {
		return cfg, errors.New("")
	}
	cfg.Server.Port = port
	// cfg.Token.JWTExpiry = os.Getenv("JWT_EXPIRY")

	cfg.Database.MaxOpenConns = 20
	cfg.Database.MaxIdleConns = 10

	return cfg, nil
}
