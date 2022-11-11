package main

import (
	"log"

	"mongogram/database"
	"mongogram/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {

	if err := database.ConnectMongo(); err != nil {
		log.Fatal(err)
	}
	app := fiber.New()

	authRoute := app.Group("/auth")

	authRoute.Post("/signup", handlers.Signup)

	app.Listen(":4000")
}
