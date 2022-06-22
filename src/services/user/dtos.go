package user

import (
	"github.com/go-playground/validator/v10"
	"rent/src/entities/users"
	"rent/src/utils/errors"
)

type LoginDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=4,alphanum"`
}

func (c LoginDto) Validate() *errors.RestError {
	v := validator.New()
	err := v.Struct(c)
	if err != nil {
		return errors.NewBadRequestError("Request error", err.Error())
	}
	return nil
}

type RegisterDto struct {
	FirstName string  `json:"first_name" validate:"required,min=2"`
	Lastname  string  `json:"last_name" validate:"required,min=2"`
	Email     string  `json:"email" validate:"required,email"`
	PhoneNo   *string `json:"phone_no" validate:"omitempty,e164"`
	Password  string  `json:"password" validate:"required,min=4,alphanum"`
	Age       *int    `json:"age" validate:"omitempty,min=18,max=100"`
}

func (c RegisterDto) Validate() *errors.RestError {
	v := validator.New()
	err := v.Struct(c)
	if err != nil {
		return errors.NewBadRequestError("Request error", err.Error())
	}
	return nil
}

type LoggedUser struct {
	User  users.User `json:"userController"`
	Token string     `json:"token"`
}
