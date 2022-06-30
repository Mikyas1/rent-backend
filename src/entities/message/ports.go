package message

import (
	"github.com/google/uuid"
	"rent/src/utils/errors"
)

type Repository interface {
	GetOrCreateConversationHead(user1, user2 uuid.UUID) (*Conversation, *errors.RestError)
	CreateMessage(mes *Message) *errors.RestError
	GetMessagesByConversationId(cId uuid.UUID) ([]Message, *errors.RestError)
	GetUserConversations(userId uuid.UUID) ([]Conversation, *errors.RestError)
}
