package QueueHandlers

import (
	"context"
	"doorProject/pkg/config"
	"doorProject/pkg/service"
	"encoding/json"
	"fmt"

	"github.com/labstack/gommon/log"
	"github.com/redis/go-redis/v9"
)

type EmailSenderHandler struct {
	redisClient  *config.RedisClient
	mailSender   service.MailSenderInterface
	subscription *redis.PubSub
}

func NewEmailSenderHandler(
	redisClient *config.RedisClient,
	mailSender service.MailSenderInterface,
) *EmailSenderHandler {
	return &EmailSenderHandler{
		redisClient: redisClient,
		mailSender:  mailSender,
	}
}

func (h *EmailSenderHandler) Handle(ctx context.Context, topic string) {
	consumerCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	h.subscription = h.redisClient.RedisClient.Subscribe(ctx, topic)
	defer h.subscription.Close()

	log.Printf("[%s] Потребитель начал прослушивание...\n", topic)

	channel := h.subscription.Channel()

	for {
		select {
		// Проверить, отменен ли основной контекст, чтобы остановить горутину
		case <-consumerCtx.Done():
			log.Printf("[%s] Потребитель прекратил прослушивание...\n", topic)
			return
			// Прослушивать входящие сообщения на канале
		case msg := <-channel:
			var messageObj interface{}
			// Десериализация полезной нагрузки сообщения
			err := json.Unmarshal([]byte(msg.Payload), &messageObj)
			if err != nil {
				log.Printf("[%s] Не удалось десериализовать сообщение: %v", topic, err)
				continue
			}
			err = h.mailSender.SendMessage(messageObj.(string))

			if err != nil {
				log.Error(err.Error())
			}

			fmt.Printf("[%s] Received message: %+v\n", topic, messageObj)
		}
	}
}
