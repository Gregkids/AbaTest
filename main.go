package main

import (
	"aba.technical.test/controller"
	"github.com/gofiber/fiber/v2"

	jwtware "github.com/gofiber/contrib/jwt"
)

func main() {
	app := fiber.New()

	// Authorization
	app.Post("/login", controller.AuthLogin)

	// Auth Test
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("aba_secret")},
	}))

	app.Get("/restricted", controller.Restricted)

	// // Playground
	// app.Get("/admin-auth")
	// app.Get("/technician-auth")
	// app.Get("/viewer-auth")

	// // Devices Endpoints
	// app.Get("/devices")
	// app.Get("/devices/:id")
	// app.Post("/devices")
	// app.Put("/devices/:id")
	// app.Delete("devices/:id")

	app.Listen(":3000")
}
