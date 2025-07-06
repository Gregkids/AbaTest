package service

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func GenerateToken(id string, userRole string) (string, error) {
	// Create the Claims
	claims := jwt.MapClaims{
		"userID": id,
		"role":   userRole,
		"exp":    time.Now().Add(time.Hour * 3).Unix(),
		"iat":    time.Now().Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("aba_secret"))
	if err != nil {
		return "", err
	}

	return t, nil
}

func MatchPassword(password, plaintext string) bool {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return false
	}

	byteHash := []byte(hashed)
	bytePlain := []byte(plaintext)

	err = bcrypt.CompareHashAndPassword(byteHash, bytePlain)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
