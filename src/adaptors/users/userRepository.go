package users

import (
	"gorm.io/gorm"
	"rent/src/entities/users"
	"rent/src/logger"
	"rent/src/utils/errors"
)

type UserRepository struct {
	db *gorm.DB
}

func (r UserRepository) AddRegularUser(user *users.User) *errors.RestError {
	result := r.db.Create(user)
	if result.Error != nil {
		logger.Error("Error while saving Regular User " + result.Error.Error())
		return errors.GormError(result.Error, "Regular User")
	}
	return nil
}

func (r UserRepository) GetRegularUser(email string) (*users.User, *errors.RestError) {
	var us users.User
	result := r.db.Where("email = ?", email).First(&us)
	if result.Error != nil {
		return nil, errors.GormError(result.Error, "User")
	}
	return &us, nil
}

func NewUserRepository(db *gorm.DB) users.Repository {
	return UserRepository{
		db: db,
	}
}
