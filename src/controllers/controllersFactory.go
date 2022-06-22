package controllers

import (
	"gorm.io/gorm"
	"rent/src/adaptors/properties"
	storage2 "rent/src/adaptors/storage"
	"rent/src/adaptors/users"
	"rent/src/controllers/propertyController"
	"rent/src/controllers/userController"
	"rent/src/services/property"
	"rent/src/services/user"
	"rent/src/storage"
)

func CreateUserController(db *gorm.DB) userController.UserController {
	service := user.NewDefaultUserService(
		users.NewUserRepository(db),
	)
	return userController.NewUserController(service)
}

func CreatePropertyController(db *gorm.DB, storage storage.Storage) propertyController.PropertyController {
	service := property.NewDefaultPropertyService(
		properties.NewPropertyRepository(db), storage2.NewMinioStorage(storage))
	return propertyController.NewPropertyController(service)
}
