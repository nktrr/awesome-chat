package message

import (
	"awesome-chat/message/internal/converter"
	"awesome-chat/message/internal/dto"
	"awesome-chat/message/internal/model"
	"awesome-chat/message/internal/repository"
	"awesome-chat/message/internal/service"
	"context"
	"github.com/google/uuid"
)

var _ service.MessageService = (*serv)(nil)

type serv struct {
	msgRepo repository.MessageRepository
}

func NewService(msgRepo repository.MessageRepository) *serv {
	return &serv{
		msgRepo: msgRepo,
	}
}

func (s serv) Create(ctx context.Context, message *model.Message) error {
	return s.msgRepo.Create(ctx, message)
}

func (s serv) Get(ctx context.Context, id uuid.UUID) (*dto.Message, error) {
	msg, err := s.msgRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return converter.ToDtoFromMessage(msg), nil
}
