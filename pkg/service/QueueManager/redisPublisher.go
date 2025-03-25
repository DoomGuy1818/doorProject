package QueueManger

import (
	"context"
	"doorProject/pkg/config"
	"encoding/json"
	"log"
)

type RedisPublisher struct {
	redisClient *config.RedisClient
}

func NewRedisPublisher(redisClient *config.RedisClient) *RedisPublisher {
	return &RedisPublisher{
		redisClient: redisClient,
	}
}

func (p *RedisPublisher) PublishMessage(ctx context.Context, message interface{}, topic string) {
	serializeMessage, err := json.Marshal(message)
	if err != nil {
		log.Printf("Error serializing message: %v", err)
	}

	err = p.redisClient.RedisClient.Publish(ctx, topic, serializeMessage).Err()

	if err != nil {
		log.Printf("Error publishing message: %v", err)
	}
}
