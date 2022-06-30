package users

import (
	"github.com/google/uuid"
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
		logger.Error("Error while getting Regular User by email " + result.Error.Error())
		return nil, errors.GormError(result.Error, "User")
	}
	return &us, nil
}

func (r UserRepository) AddRenterProfile(userId uuid.UUID, rp users.RenterProfile) (*users.User, *errors.RestError) {
	var user users.User
	result := r.db.Model(&user).
		Where("id = ?", userId.String()).
		Association("RenterProfile").
		Append(&rp)
	if result != nil {
		logger.Error("Error while saving renter profile " + result.Error())
		return nil, errors.GormError(result, "Renter Profile")
	}

	result2 := r.db.Where("id = ?", userId.String()).First(&user)
	if result2.Error != nil {
		logger.Error("Error while getting Regular User by id " + result2.Error.Error())
		return nil, errors.GormError(result2.Error, "User")
	}

	user.RenterProfileId = &rp.UUID
	*user.RenterProfile = rp
	return &user, nil
}

func (r UserRepository) GetRenterProfile(userId uuid.UUID) (*users.RenterProfile, *errors.RestError) {
	var us users.User
	result := r.db.Where("id = ?", userId.String()).Preload("RenterProfile").First(&us)
	if result.Error != nil {
		logger.Error("Error while getting Regular User by id " + result.Error.Error())
		return nil, errors.GormError(result.Error, "User")
	}

	return us.RenterProfile, nil
}

func NewUserRepository(db *gorm.DB) users.Repository {
	return UserRepository{
		db: db,
	}
}
