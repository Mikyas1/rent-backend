package message

import (
	"github.com/google/uuid"
	"rent/src/entities/message"
	"rent/src/utils/errors"
)

type MessageService interface {
	GetConversationHeads(userId uuid.UUID) ([]message.Conversation, *errors.RestError)
	GetMessages(cId uuid.UUID) ([]message.Message, *errors.RestError)
	SendMessage(dto SendMessageDto) *errors.RestError
}
