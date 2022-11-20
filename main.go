package main

import (
	"log"

	"mongogram/database"
	"mongogram/routers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {

	if err := database.ConnectMongo(); err != nil {
		log.Fatal(err)
	}
	app := fiber.New()
	app.Use(recover.New())
	routers.SetupRoute(app)

	app.Listen(":4000")
}
