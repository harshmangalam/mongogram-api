package router

import (
	"mongogram/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoute(app *fiber.App) {
	authRoute := app.Group("/auth")

	authRoute.Post("/signup", handler.Signup)
}
