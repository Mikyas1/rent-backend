package message

import (
	"github.com/google/uuid"
	"rent/src/entities"
	"rent/src/entities/users"
)

type Conversation struct {
	entities.Model
	Participants string `json:"participants" gorm:"unique;index; not null"`
}

func (c *Conversation) CreateConversation(user1, user2 uuid.UUID) {
	c.UUID = uuid.New()
	c.Participants = user1.String() + " " + user2.String()
}

type Message struct {
	entities.Model
	Message        string       `json:"message"`
	ConversationId uuid.UUID    `json:"conversation_id" gorm:"type:char(36);index;not null"`
	Conversation   Conversation `json:"conversation" gorm:"foreignKey:ConversationId;references:UUID"`

	SenderId uuid.UUID  `json:"sender_id" gorm:"type:char(36);index;not null"`
	Sender   users.User `json:"sender" gorm:"foreignKey:SenderId;references:UUID"`

	ReceiverId uuid.UUID  `json:"receiver_id" gorm:"type:char(36);index;not null"`
	Receiver   users.User `json:"receiver" gorm:"foreignKey:ReceiverId;references:UUID"`
}
