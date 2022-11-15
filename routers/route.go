package routers

import (
	"mongogram/handlers/auth"
	"mongogram/handlers/search"
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
	userRoute.Get("/:userId", user.GetUser)
	userRoute.Post("/:userId/friendship", middlewares.Protected(), user.Friendship)

	// search route
	searchRoute := app.Group("/search")
	searchRoute.Get("/", middlewares.Protected(), search.AtlasSearch)
	searchRoute.Get("/recent", middlewares.Protected(), search.GetRecentSearch)
	searchRoute.Delete("/:searchId", middlewares.Protected(), search.DeleteRecentSearch)

}
