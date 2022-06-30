package users

import (
	"github.com/google/uuid"
	"rent/src/utils/errors"
)

type Repository interface {
	AddRegularUser(*User) *errors.RestError
	GetRegularUser(email string) (*User, *errors.RestError)
	AddRenterProfile(userId uuid.UUID, renterProfile RenterProfile) (*User, *errors.RestError)
	GetRenterProfile(userId uuid.UUID) (*RenterProfile, *errors.RestError)
}
