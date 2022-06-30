package message

import (
	"github.com/google/uuid"
	"rent/src/entities"
	"rent/src/entities/users"
	"time"
)

type Conversation struct {
	entities.Model
	LatestMessage time.Time `json:"latest_message"`
}

type Message struct {
	entities.Model
	Message        string     `json:"message"`
	ConversationId uuid.UUID  `json:"conversation_id" gorm:"type:char(36);index;not null"`
	Conversation   users.User `json:"conversation" gorm:"foreignKey:ConversationId;references:UUID"`
}
