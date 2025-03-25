package service

import "context"

type MessageHandlerInterface interface {
	Handle(ctx context.Context, topic string)
}
