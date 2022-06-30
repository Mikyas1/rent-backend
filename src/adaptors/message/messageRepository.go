package message

import (
	errors2 "errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"rent/src/entities/message"
	"rent/src/logger"
	"rent/src/utils/errors"
)

type MessageRepository struct {
	db *gorm.DB
}

func (r MessageRepository) GetOrCreateConversationHead(user1, user2 uuid.UUID) (*message.Conversation, *errors.RestError) {
	var con message.Conversation
	result := r.db.Where("participants LIKE ?", user1.String()).
		Where("participants LIKE ?", user2.String()).
		First(&con)
	if result.Error != nil {
		if errors2.Is(result.Error, gorm.ErrRecordNotFound) {
			//	create conversation
			con.CreateConversation(user1, user2)

			result = r.db.Create(&con)
			if result.Error != nil {
				logger.Error("Error while saving Conversation " + result.Error.Error())
				return nil, errors.GormError(result.Error, "Conversation")
			}
			return &con, nil
		} else {
			logger.Error("Error when getting Conversation by user1 and user2 " + result.Error.Error())
			return nil, errors.GormError(result.Error, "Conversation")
		}
	}
	return &con, nil
}

func (r MessageRepository) CreateMessage(mes *message.Message) *errors.RestError {
	result := r.db.Create(mes)
	if result.Error != nil {
		logger.Error("Error while saving Message " + result.Error.Error())
		return errors.GormError(result.Error, "Message")
	}
	return nil
}

func (r MessageRepository) GetMessagesByConversationId(cId uuid.UUID) ([]message.Message, *errors.RestError) {
	var msgs []message.Message
	result := r.db.Where("conversation_id = ?", cId.String()).
		Find(&msgs)
	if result.Error != nil {
		logger.Error("Error when getting Messages by conversation Id" + result.Error.Error())
		return nil, errors.GormError(result.Error, "Message")
	}
	return msgs, nil
}

func (r MessageRepository) GetUserConversations(userId uuid.UUID) ([]message.Conversation, *errors.RestError) {
	var cons []message.Conversation
	result := r.db.Where("participants LIKE ?", userId.String()).
		Find(&cons)
	if result.Error != nil {
		logger.Error("Error when getting Conversations by userId" + result.Error.Error())
		return nil, errors.GormError(result.Error, "Conversation")
	}
	return cons, nil
}

func NewMessageRepository(db *gorm.DB) message.Repository {
	return MessageRepository{
		db: db,
	}
}
