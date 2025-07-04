package main

import (
	"github.com/gofiber/fiber/v2"

	jwtware "github.com/gofiber/contrib/jwt"
)

func main() {
	app := fiber.New()

	// Authorization
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	}))
	app.Post("/login")

	// Auth Test
	app.Get("/admin-auth")
	app.Get("/technician-auth")
	app.Get("/viewer-aut")

	// Devices Endpoints
	app.Get("/devices")
	app.Get("/devices/:id")
	app.Post("/devices")
	app.Put("/devices/:id")
	app.Delete("devices/:id")

	app.Listen(":3000")
}
