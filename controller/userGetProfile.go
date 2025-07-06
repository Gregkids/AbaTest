package controller

import (
	"database/sql"

	"aba.technical.test/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func UserGetProfile(c *fiber.Ctx) error {
	// Connecting to Database
	db, err := sql.Open("postgres", "host=localhost user=postgres password=darageta dbname=local_test sslmode=disable")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": fiber.ErrInternalServerError,
			"desc":   err.Error(),
		})
	}

	// Parse Database
	r := repository.UserSQL{DB: db}

	// Declare Request
	user := c.Locals("user").(*jwt.Token)
	claim := user.Claims.(jwt.MapClaims)
	req := claim["userID"].(string)

	// Execute Method
	ret, err := r.GetUser(req)
	if err != nil {
		return c.Status(422).JSON(fiber.Map{
			"status": fiber.ErrUnprocessableEntity,
			"desc":   err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"code":   fiber.StatusOK,
		"device": ret,
	})
}
