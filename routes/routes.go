package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	jwtware "github.com/gofiber/jwt/v3"
	"os"
)

func Setup(app *fiber.App) {
	// Non authenticated routes
	app.Get("/dashboard", monitor.New())

	// Authenticate Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("jwtsecret")),
	}))

	// Authenticated Routes

}
