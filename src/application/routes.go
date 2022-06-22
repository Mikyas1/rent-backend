package application

import (
	"github.com/gofiber/fiber/v2"
	"rent/src/controllers/routes"
)

func RegisterApi(app *fiber.App) {
	routes.UserAuthRoutes(app, Db.ProvideDb())
	routes.PropertyRoutesProtected(app, Db.ProvideDb(), Storage)
}
