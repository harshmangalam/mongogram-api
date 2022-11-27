package routers

import (
	"mongogram/handlers/accounts"
	"mongogram/handlers/auth"
	"mongogram/handlers/media"
	"mongogram/handlers/post"
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
	authRoute.Post("/login-with-github", auth.LoginWithGithub)

	// user route
	userRoute := app.Group("/user")
	userRoute.Get("/me", middlewares.Protected(), user.GetMe)
	userRoute.Get("/:userId", user.GetUser)
	userRoute.Post("/:userId/friendship", middlewares.Protected(), user.Friendship)

	// search route
	searchRoute := app.Group("/search")
	searchRoute.Get("/", middlewares.Protected(), search.AtlasSearch)
	searchRoute.Get("/recent", middlewares.Protected(), search.GetRecentSearch)
	searchRoute.Delete("/recent", middlewares.Protected(), search.DeleteRecentSearchs)
	searchRoute.Delete("/recent/:searchId", middlewares.Protected(), search.DeleteRecentSearch)

	// account route
	accountRoute := app.Group("/account")
	accountRoute.Put("/edit", middlewares.Protected(), accounts.EditAccount)
	accountRoute.Patch("/change_password", middlewares.Protected(), accounts.ChangePassword)
	accountRoute.Post("/reset_password", accounts.ResetPassword)

	//posts
	postRoute := app.Group("/posts")
	postRoute.Post("/", middlewares.Protected(), post.CreatePost)
	postRoute.Get("/", post.GetPosts)
	postRoute.Get("/:postId", post.GetPost)

	// media
	mediaRoute := app.Group("/media")
	mediaRoute.Post("/upload", middlewares.Protected(), media.Upload)
	mediaRoute.Get("/", media.GetMedia)

}
