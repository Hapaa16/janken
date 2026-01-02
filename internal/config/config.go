package config

import (
	"os"

	"github.com/Hapaa16/janken/internal/infra/db"
	"github.com/Hapaa16/janken/internal/infra/redis"
)

type Config struct {
	Env      string
	Port     string
	DB       db.Config
	Redis    redis.Config
	ServerId string
}

func Load() *Config {
	return &Config{
		Env:  getEnv("ENV", "local"),
		Port: getEnv("PORT", "8080"),
		DB: db.Config{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "janken-service"),
			Password: getEnv("DB_PASSWORD", "pass123"),
			Name:     getEnv("DB_NAME", "janken-service-db"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		Redis: redis.Config{
			Addr:     getEnv("REDIS_ADDR", "localhost:6379"),
			Password: getEnv("REDIS_PASSWORD", ""),
			DB:       1,
			Protocol: 6,
		},
		ServerId: "1",
	}
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
