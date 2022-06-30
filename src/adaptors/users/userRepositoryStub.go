package users

import (
	"github.com/google/uuid"
	"rent/src/entities/users"
	"rent/src/utils/errors"
)

type UserRepositoryStub struct {
	SampleUser users.User
}

func (r UserRepositoryStub) AddRegularUser(user *users.User) *errors.RestError {
	return nil
}

func (r UserRepositoryStub) GetRegularUser(email string) (*users.User, *errors.RestError) {
	r.SampleUser.Email = email
	return &r.SampleUser, nil
}

func (r UserRepositoryStub) AddRenterProfile(userId uuid.UUID, dto users.RenterProfile) (*users.User, *errors.RestError) {
	return nil, nil
}

func (r UserRepositoryStub) GetRenterProfile(userId uuid.UUID) (*users.RenterProfile, *errors.RestError) {
	return nil, nil
}

func NewUserRepositoryStub(user users.User) users.Repository {
	return UserRepositoryStub{
		SampleUser: user,
	}
}
