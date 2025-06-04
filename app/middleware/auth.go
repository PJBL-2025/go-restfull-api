package middleware

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"strings"
)

var secretKey = []byte("secret")

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {

		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized 1"})
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized 2"})
		}

		token, err := jwt.Parse(tokenParts[1], func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return secretKey, nil
		})

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized 3"})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized 4"})
		}

		c.Locals("userId", int(claims["id"].(float64)))
		c.Locals("role", claims["role"])

		return c.Next()
	}
}
