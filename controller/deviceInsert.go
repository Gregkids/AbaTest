package controller

import (
	"database/sql"

	"aba.technical.test/models"
	"aba.technical.test/repository"
	"github.com/gofiber/fiber/v2"
)

func DeviceAdd(c *fiber.Ctx) error {
	// Connecting to Database
	db, err := sql.Open("postgres", "host=localhost user=postgres password=darageta dbname=local_test sslmode=disable")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": fiber.ErrInternalServerError,
			"desc":   err.Error(),
		})
	}

	// Parse Database
	r := repository.DeviceSQL{DB: db}

	// Declare Requests
	req := new(models.DeviceReq)
	err = c.BodyParser(req)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": fiber.ErrBadRequest,
			"desc":   err.Error(),
		})
	}

	// Execute Method
	err = r.InsertDevice(req)
	if err != nil {
		return c.Status(422).JSON(fiber.Map{
			"status": fiber.ErrUnprocessableEntity,
			"desc":   err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "Device Added",
	})
}
