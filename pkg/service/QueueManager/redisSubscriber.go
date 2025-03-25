package QueueManger

import (
	"context"
	"doorProject/pkg/config"
	"doorProject/pkg/service"

	"github.com/labstack/gommon/log"
)

type RedisSubscriber struct {
	client   *config.RedisClient
	handlers map[string]service.MessageHandlerInterface
}

func NewRedisSubscriber(
	client *config.RedisClient,
	handlers map[string]service.MessageHandlerInterface,
) *RedisSubscriber {
	return &RedisSubscriber{client: client, handlers: handlers}
}

func (s *RedisSubscriber) RegisterHandler(topic string, handler service.MessageHandlerInterface) {
	if handler == nil {
		log.Errorf("Proccess handler is nil")
	}

	s.handlers[topic] = handler
}

func (s *RedisSubscriber) ConsumeMessages(ctx context.Context, topics []string) {
	for _, queueName := range topics {
		handler, exists := s.handlers[queueName]
		if !exists {
			log.Printf("[%s] Неподдерживаемый тип сообщения\n", queueName)
			continue
		}
		go handler.Handle(ctx, queueName)
	}
}
