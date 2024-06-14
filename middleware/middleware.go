package middleware

import (
	"fmt"
	"gts-dry/model"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthMiddleware(secret string) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		// Get token from request header
		tokenString := c.Get("Authorization")
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// // Parse the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if token.Method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("invalid signing method")
			}
			return []byte(secret), nil
		})

		// If token is invalid or expired
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(model.ErrorResponse{
				Success: false,
				Code:    fiber.StatusUnauthorized,
				Message: "Unauthorized",
				Error:   err.Error(),
			})
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Locals("id", claims["id"])
			c.Next()
			return nil
		} else {
			return c.Status(fiber.StatusUnauthorized).JSON(model.ErrorResponse{
				Success: false,
				Code:    fiber.StatusUnauthorized,
				Message: "Unauthorized",
				Error:   "Unauthorized",
			})
		}
	}
}
