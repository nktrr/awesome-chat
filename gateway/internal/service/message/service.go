package message

import (
	"awesome-chat/gateway/internal/dto"
	"awesome-chat/gateway/internal/kafka"
	"awesome-chat/gateway/internal/model"
	"awesome-chat/gateway/internal/service"
	"context"
	"github.com/google/uuid"
)

var _ service.MessageService = (*serv)(nil)

type serv struct {
	msgProducer *kafka.Kafka
}

func NewService(msgProducer *kafka.Kafka) *serv {
	return &serv{
		msgProducer: msgProducer,
	}
}

func (s serv) Create(ctx context.Context, message *model.Message) error {
	return s.msgProducer.Write(message)
}

func (s serv) Get(ctx context.Context, id uuid.UUID) (*dto.Message, error) {
	return nil, nil
	//msg, err := s.msgRepo.Get(ctx, id)
	//if err != nil {
	//	return nil, err
	//}
	//return converter.ToDtoFromMessage(msg), nil
}
