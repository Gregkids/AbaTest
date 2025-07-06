package controller

import (
	"database/sql"

	"aba.technical.test/models"
	"aba.technical.test/repository"
	"github.com/gofiber/fiber/v2"
)

func DeviceUpdate(c *fiber.Ctx) error {
	// Connecting to Database
	db, err := sql.Open("postgres", "host=localhost user=postgres password=darageta dbname=local_test sslmode=disable")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"code":  fiber.StatusInternalServerError,
			"error": fiber.ErrInternalServerError,
			"msg":   err.Error(),
		})
	}

	// Parse Database
	r := repository.DeviceSQL{DB: db}

	// Declare Requests
	reqDevice := new(models.DeviceReq)
	reqID := c.Params("id")
	err = c.BodyParser(reqDevice)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"code":  fiber.StatusBadRequest,
			"error": fiber.ErrBadRequest,
			"msg":   err.Error(),
		})
	}

	// Execute Method
	err = r.UpdateDevice(reqID, reqDevice)
	if err != nil {
		return c.Status(422).JSON(fiber.Map{
			"code":  fiber.StatusUnprocessableEntity,
			"error": fiber.ErrUnprocessableEntity,
			"msg":   err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"code": fiber.StatusOK,
		"msg":  "Device Updated",
	})
}
