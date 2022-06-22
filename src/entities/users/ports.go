package users

import "rent/src/utils/errors"

type Repository interface {
	AddRegularUser(*User) *errors.RestError
	GetRegularUser(email string) (*User, *errors.RestError)
}
