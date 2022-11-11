package router

import (
	"mongogram/handler/auth"

	"github.com/gofiber/fiber/v2"
)

func SetupRoute(app *fiber.App) {
	authRoute := app.Group("/auth")

	authRoute.Post("/signup", auth.Signup)
	authRoute.Post("/login", auth.Login)
}
