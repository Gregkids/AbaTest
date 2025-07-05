package service

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type Token struct {
	UserId string `json:"uid"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

func GenerateToken(id string, userRole string) (string, error) {
	// Create the Claims
	claims := Token{
		id,
		userRole,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return t, nil
}

func MatchPassword(hashed, plaintext string) bool {
	byteHash := []byte(hashed)
	bytePlain := []byte(plaintext)

	err := bcrypt.CompareHashAndPassword(byteHash, bytePlain)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
