package controller

import (
	"database/sql"

	"aba.technical.test/repository"
	"github.com/gofiber/fiber/v2"
)

func DeviceGetAll(c *fiber.Ctx) error {
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

	// Execute Method
	ret, err := r.GetAllDevice()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"code":  fiber.StatusNotFound,
			"error": fiber.ErrNotFound,
			"msg":   err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"Devices": ret,
	})
}
