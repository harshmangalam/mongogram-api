package routers

import (
	"mongogram/handlers/auth"
	"mongogram/handlers/user"
	"mongogram/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetupRoute(app *fiber.App) {

	// auth route
	authRoute := app.Group("/auth")

	authRoute.Post("/signup", auth.Signup)
	authRoute.Post("/login", auth.Login)

	// user route

	userRoute := app.Group("/user")
	userRoute.Get("/me", middlewares.Protected(), user.GetMe)
}
