package controller

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthLogin(c *fiber.Ctx) error {
	// // Connecting to Database
	// db, err := sql.Open("postgres", "host=localhost user=postgres password=darageta dbname=postgres sslmode=disable")
	// if err != nil {
	// 	return c.Status(500).JSON(fiber.Map{
	// 		"code":   500,
	// 		"status": "SQL Error",
	// 		"msg":    err.Error(),
	// 	})
	// }

	// // Parse Database
	// r := repository.UserSQL{DB: db}

	user := c.FormValue("user")
	pass := c.FormValue("pass")

	// Throws Unauthorized error
	if user != "john" || pass != "doe" {
		fmt.Print("Error")
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"name":  "John Doe",
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	fmt.Print(t)
	return nil
}
