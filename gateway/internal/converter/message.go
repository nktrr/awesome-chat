package converter

import (
	"awesome-chat/gateway/internal/dto"
	"awesome-chat/gateway/internal/model"
)

func ToDtoFromMessage(message *model.Message) *dto.Message {
	return &dto.Message{
		ID:   message.ID.String(),
		Text: message.Text,
	}
}

// TODO [18.09.2024] коряво, надо разобраться какая модель на транспортном уровне и как вниз прокидывать
func ToMessageFromDto(message *dto.Message) *model.Message {
	return &model.Message{
		// MustParse wtih "" == panic
		//ID:   uuid.MustParse(message.ID),
		Text: message.Text,
	}
}
