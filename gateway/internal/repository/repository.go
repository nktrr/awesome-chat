package repository

import (
	"awesome-chat/gateway/internal/model"
	"context"
)

type MessageRepository interface {
	Create(ctx context.Context, message model.Message) error
	Get(ctx context.Context) error
}
