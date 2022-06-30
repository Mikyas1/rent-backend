package message

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"rent/src/utils/errors"
)

type SendMessageDto struct {
	Message    string    `json:"message" validate:"required,min=1"`
	SenderId   uuid.UUID `json:"sender_id" validate:"required"`
	ReceiverId uuid.UUID `json:"receiver_id" validate:"required"`
}

func (d SendMessageDto) Validate() *errors.RestError {
	v := validator.New()
	err := v.Struct(d)
	if err != nil {
		return errors.NewBadRequestError("Request error", err.Error())
	}
	return nil
}
