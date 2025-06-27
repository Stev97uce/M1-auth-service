package config

import (
	"os"
)

type Config struct {
	RedisHost  string
	RedisPort  string
	RedisPass  string
	SessionTTL string
}

func LoadConfig() *Config {
	return &Config{
		RedisHost:  getEnv("REDIS_HOST", "localhost"),
		RedisPort:  getEnv("REDIS_PORT", "6379"),
		RedisPass:  getEnv("REDIS_PASSWORD", ""),
		SessionTTL: getEnv("SESSION_TTL", "3600"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
