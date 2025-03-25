package service

import "context"

type MessageConsumerInterface interface {
	ConsumeMessages(ctx context.Context, topics []string)
}
