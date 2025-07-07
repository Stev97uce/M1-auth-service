package config

import (
	"os"
)

type Config struct {
	RedisHost             string
	RedisPort             string
	RedisPass             string
	SessionTTL            string
	UserProfileServiceURL string
}

func LoadConfig() *Config {
	return &Config{
		RedisHost:             getEnv("REDIS_HOST", "18.232.41.255"),
		RedisPort:             getEnv("REDIS_PORT", "6379"),
		RedisPass:             getEnv("REDIS_PASSWORD", ""),
		SessionTTL:            getEnv("SESSION_TTL", "3600"),
		UserProfileServiceURL: getEnv("USER_PROFILE_SERVICE_URL", "http://34.193.149.183:8000"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
