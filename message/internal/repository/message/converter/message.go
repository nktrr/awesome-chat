package converter

import (
	"awesome-chat/message/internal/model"
	modelRepo "awesome-chat/message/internal/repository/message/model"
	"github.com/google/uuid"
)

func ToMessageFromRepo(message *modelRepo.Message) *model.Message {
	return &model.Message{
		ID:   uuid.UUID(message.ID),
		Text: message.Text,
	}
}
