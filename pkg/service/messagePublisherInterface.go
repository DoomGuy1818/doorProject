package service

import "context"

type MessagePublisherInterface interface {
	PublishMessage(ctx context.Context, message interface{}, topic string)
}
