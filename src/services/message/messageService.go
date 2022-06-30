package message

import (
	"github.com/google/uuid"
	"rent/src/entities/message"
	"rent/src/utils/errors"
)

type DefaultMessageService struct {
	repo message.Repository
}

func (s DefaultMessageService) GetConversationHeads(userId uuid.UUID) ([]message.Conversation, *errors.RestError) {
	return s.repo.GetUserConversations(userId)
}

func (s DefaultMessageService) GetMessages(cId uuid.UUID) ([]message.Message, *errors.RestError) {
	return s.repo.GetMessagesByConversationId(cId)
}

func (s DefaultMessageService) SendMessage(dto SendMessageDto) *errors.RestError {
	return nil
}

func NewMessageService(repo message.Repository) MessageService {
	return DefaultMessageService{
		repo: repo,
	}
}
