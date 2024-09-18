package message

import (
	"awesome-chat/gateway/internal/model"
	"awesome-chat/gateway/internal/repository"
	"awesome-chat/gateway/internal/service"
	"context"
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

func (s serv) Create(ctx context.Context, message model.Message) error {
	s.msgRepo.Create(ctx, message)
	return nil
}

func (s serv) Get(ctx context.Context) error {
	return nil
}
