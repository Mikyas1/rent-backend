package application

import (
	"github.com/gofiber/fiber/v2"
	"rent/src/controllers/routes"
)

func RegisterApi(app *fiber.App) {
	routes.UserAuthRoutes(app, Db.ProvideDb())
	routes.PropertyRoutes(app, Db.ProvideDb(), Storage)
	routes.ObjectStorageRoutes(app, Storage)
	routes.RentRoutes(app, Db.ProvideDb())
	routes.AdminRoutes(app, Db.ProvideDb())
}
