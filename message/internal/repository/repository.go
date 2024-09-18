package repository

import (
	"awesome-chat/message/internal/model"
	"context"
	"github.com/google/uuid"
)

type MessageRepository interface {
	Create(ctx context.Context, message *model.Message) error
	Get(ctx context.Context, id uuid.UUID) (*model.Message, error)
}
