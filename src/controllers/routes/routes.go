package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"rent/src/controllers"
	"rent/src/middleware"
	"rent/src/storage"
)

func UserAuthRoutes(api fiber.Router, db *gorm.DB) {
	userController := controllers.CreateUserController(db)

	auth := api.Group("auth")

	auth.Post("register", userController.RegisterUser)
	auth.Post("login", userController.LoginUser)
}

func PropertyRoutesProtected(api fiber.Router, db *gorm.DB, storage storage.Storage) {
	propertyController := controllers.CreatePropertyController(db, storage)

	property := api.Group("property")

	property.Use(middleware.ValidateUser)
	property.Post("add-property", propertyController.AddProperty)
}
