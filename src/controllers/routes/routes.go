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

	auth.Use(middleware.ValidateUser)
	auth.Post("renter-profile", userController.CreateRenterProfile)
	auth.Get("renter-profile", userController.GetRenterProfile)
}

func PropertyRoutes(api fiber.Router, db *gorm.DB, storage storage.Storage) {
	propertyController := controllers.CreatePropertyController(db, storage)

	property := api.Group("property", middleware.GeneratePaginationFromRequest)
	property.Get("property/:id", propertyController.GetPropertyDetail)
	property.Post("top-properties", propertyController.GetTopProperties)
	property.Get("property-options", propertyController.GetPropertyOptions)

	property.Use(middleware.ValidateUser)
	property.Post("add-property", propertyController.AddProperty)
	property.Get("my-properties", propertyController.GetOwnerProperties)
	property.Delete(":id", propertyController.RemoveProperty)
}

func ObjectStorageRoutes(api fiber.Router, storage storage.Storage) {
	objectStorageHandler := controllers.CreateObjectStorageHandler(storage)

	media := api.Group("media")
	media.Get("/:object_name", objectStorageHandler.File)
}

func RentRoutes(api fiber.Router, db *gorm.DB) {
	rentController := controllers.CreateRentController(db)
	rent := api.Group("rent", middleware.ValidateUser)

	rent.Get("payment-types", rentController.GetPaymentTypes)
	rent.Post("create-rent-request", rentController.CreateRentRequest)
	rent.Get("get-rent-request-as-owner", rentController.GetRentRequestAsOwner)
	rent.Get("get-rent-request-as-renter", rentController.GetRentRequestAsRenter)
	rent.Get("rent-request-detail/:id", rentController.RentRequestDetail)
	rent.Get("reject-rent-request/:id", rentController.RejectRentRequest)
	rent.Get("accept-rent-request/:id", rentController.AcceptRentRequest)
	rent.Get("rents-as-owner", rentController.AllRentsAsOwner)
	rent.Get("rents-as-renter", rentController.AllRentsAsRenter)
	rent.Get("rent-detail/:id", rentController.GetRentDetail)
}

func AdminRoutes(api fiber.Router, db *gorm.DB) {
	adminController := controllers.CreateAdminController(db)
	admin := api.Group("admin")

	admin.Get("statistics", adminController.GetStatistics)
	admin.Get("pending-properties", adminController.GetPendingProperties)
	admin.Get("approve-property/:id", adminController.ApproveProperty)
	admin.Get("reject-property/:id", adminController.RejectProperty)
}

func MessageRoutes(api fiber.Router, db *gorm.DB) {
	messageController := controllers.CreateMessageController(db)
	message := api.Group("message", middleware.ValidateUser)

	message.Get("conversations", messageController.GetConversations)
	message.Get("get-messages/:id", messageController.GetMessages)
	message.Post("send-message", messageController.SendMessage)
}
