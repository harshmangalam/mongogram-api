package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

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

	go func() {
		if err := app.Listen(":4000"); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel
	<-c                                             // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down...")
	_ = app.Shutdown()
	fmt.Println("Running cleanup tasks...")

	// Add your cleanup tasks here...

	fmt.Println("Fiber was successful shutdown.")
}
