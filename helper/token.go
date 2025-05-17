package helper

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secret")

func TokenSeeder() {
	claims := jwt.MapClaims{
		"id":    2,
		"role":  "user",
		"email": "jrocket@example.com",
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println("Gagal generate token:", err)
		return
	}

	fmt.Println("Token JWT:", tokenString)
}
