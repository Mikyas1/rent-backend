package user

import (
	"github.com/google/uuid"
	"rent/src/entities/users"
	"rent/src/utils/errors"
)

type UserService interface {
	RegisterUser(dto RegisterDto) (*LoggedUser, *errors.RestError)
	LoginUser(dto LoginDto) (*LoggedUser, *errors.RestError)
	AddRenterProfile(userId uuid.UUID, dto RenterProfileDto) (*users.RenterProfile, *errors.RestError)
	GetRenterProfile(userId uuid.UUID) (*users.RenterProfile, *errors.RestError)
}
