package main

import (
	"log"

	"mongogram/database"
	"mongogram/routers"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	if err := database.ConnectMongo(); err != nil {
		log.Fatal(err)
	}
	app := fiber.New()
	routers.SetupRoute(app)

	app.Listen(":4000")
}
