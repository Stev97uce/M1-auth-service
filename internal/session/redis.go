package session

import (
	"context"
	"fmt"
	"time"

	"auth-service/config"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type RedisClient struct {
	Client *redis.Client
	TTL    time.Duration
}

func NewRedisClient(cfg *config.Config) *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
		Password: cfg.RedisPass,
		DB:       0,
	})

	ttl, _ := time.ParseDuration(cfg.SessionTTL + "s")

	return &RedisClient{
		Client: client,
		TTL:    ttl,
	}
}

func (r *RedisClient) SetSession(token, userID string) error {
	return r.Client.Set(ctx, token, userID, r.TTL).Err()
}

func (r *RedisClient) DeleteSession(token string) error {
	return r.Client.Del(ctx, token).Err()
}

func (r *RedisClient) ValidateSession(token string) (string, error) {
	return r.Client.Get(ctx, token).Result()
}
