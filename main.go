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

	// User Profile Endpoints
	app.Get("/user", controller.UserGetProfile)

	// Devices Endpoints
	app.Get("/devices", controller.DeviceGetAll)
	app.Get("/devices/:id", controller.DeviceGetOne)
	app.Post("/devices", controller.DeviceAdd)
	app.Put("/devices/:id", controller.DeviceUpdate)
	app.Delete("devices/:id", controller.DeviceDelete)

	app.Listen(":3000")
}
