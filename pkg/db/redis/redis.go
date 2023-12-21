package redis

import (
	"context"
	"github.com/blazee5/tgbot/config"
	"github.com/redis/go-redis/v9"
	"log"
)

func NewRedis(cfg *config.Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Host,
		Password: cfg.Redis.Password,
		DB:       0,
	})

	err := client.Ping(context.Background()).Err()

	if err != nil {
		log.Fatalf("error while connect to redis: %v", err)
	}

	return client
}
