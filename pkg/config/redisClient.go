package config

import (
	"errors"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	RedisClient *redis.Client
}

func NewRedisClient(dsn string) *RedisClient {
	var client = redis.NewClient(
		&redis.Options{
			Addr: dsn,
		},
	)

	if client == nil {
		errors.New("cannot load redis")
	}

	return &RedisClient{
		RedisClient: client,
	}
}
