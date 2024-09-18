package service

import (
	"awesome-chat/gateway/internal/model"
	"context"
)

type MessageService interface {
	Create(ctx context.Context, message model.Message) error
	Get(ctx context.Context) error
}
