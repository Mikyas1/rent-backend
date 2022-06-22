package user

import (
	"rent/src/utils/errors"
)

type UserService interface {
	RegisterUser(dto RegisterDto) (*LoggedUser, *errors.RestError)
	LoginUser(dto LoginDto) (*LoggedUser, *errors.RestError)
}
