package main

import (
	"log"

	"mongogram/database"
	"mongogram/routers"

	"github.com/gofiber/fiber/v2"
)

func main() {

	if err := database.ConnectMongo(); err != nil {
		log.Fatal(err)
	}
	app := fiber.New()
	routers.SetupRoute(app)

	app.Listen(":4000")
}
