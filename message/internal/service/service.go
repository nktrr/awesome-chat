package service

import (
	"awesome-chat/message/internal/dto"
	"awesome-chat/message/internal/model"
	"context"
	"github.com/google/uuid"
)

type MessageService interface {
	Create(ctx context.Context, message *model.Message) error
	Get(ctx context.Context, id uuid.UUID) (*dto.Message, error)
}
