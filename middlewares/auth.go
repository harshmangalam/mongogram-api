package middlewares

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Protected protect routes
func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:     []byte("SECRET"),
		ErrorHandler:   jwtError,
		SuccessHandler: success,
		ContextKey:     "jwtToken",
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		c.Next()
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
}

func success(c *fiber.Ctx) error {
	token := c.Locals("jwtToken").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	strId := claims["userId"].(string)
	userId, _ := primitive.ObjectIDFromHex(strId)

	c.Locals("userId", userId)
	return c.Next()
}
