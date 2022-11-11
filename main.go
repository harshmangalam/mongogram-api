package main

import (
	"log"

	"mongogram/auth"
	"mongogram/database"

	"github.com/gofiber/fiber/v2"
)

func main() {

	if err := database.ConnectMongo(); err != nil {
		log.Fatal(err)
	}
	app := fiber.New()

	authRoute := app.Group("/auth")

	authRoute.Post("/signup", auth.Signup)

	app.Listen(":4000")
}
