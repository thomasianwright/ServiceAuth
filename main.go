package main

import (
	"github.com/gofiber/fiber/v2"
	"os"
	"serviceauth/database"
	"serviceauth/routes"
)

func main() {
	app := fiber.New()

	database.Setup()

	routes.Setup(app)

	port := os.Getenv("port")

	if port == "" {
		port = ":3010"
	}

	app.Listen(port)
}
