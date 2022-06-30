package controllers

import (
	"gorm.io/gorm"
	message2 "rent/src/adaptors/message"
	properties2 "rent/src/adaptors/properties"
	rent2 "rent/src/adaptors/rent"
	storage2 "rent/src/adaptors/storage"
	"rent/src/adaptors/users"
	"rent/src/controllers/adminController"
	"rent/src/controllers/messageController"
	"rent/src/controllers/objectStorageProxy"
	"rent/src/controllers/propertyController"
	"rent/src/controllers/rentControllers"
	"rent/src/controllers/userController"
	"rent/src/services/admin"
	"rent/src/services/message"
	"rent/src/services/objectStorage"
	"rent/src/services/property"
	"rent/src/services/rent"
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
		properties2.NewPropertyRepository(db), storage2.NewMinioStorage(storage))
	return propertyController.NewPropertyController(service)
}

func CreateObjectStorageHandler(storage storage.Storage) objectStorageProxy.ObjectStorageHandler {
	service := objectStorage.NewDefaultObjectStorageService(storage)
	return objectStorageProxy.NewObjectStorageHandler(service)
}

func CreateRentController(db *gorm.DB) rentControllers.RentController {
	service := rent.NewDefaultRentService(rent2.NewRentRepository(db))
	return rentControllers.NewRentController(service)
}

func CreateAdminController(db *gorm.DB) adminController.AdminController {
	adminService := admin.NewDefaultAdminService(properties2.NewPropertyRepository(db), rent2.NewRentRepository(db))
	return adminController.NewAdminController(adminService)
}

func CreateMessageController(db *gorm.DB) messageController.MessageController {
	service := message.NewMessageService(message2.NewMessageRepository(db))
	return messageController.NewMessageController(service)
}
