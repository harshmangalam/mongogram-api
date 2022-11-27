package utils

import "github.com/gofiber/fiber/v2"

type StatusText string

const (
	Success StatusText = "success"
	Error   StatusText = "error"
)

func CustomResponse(c *fiber.Ctx, statusText string, statusCode int, message string, data any) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"status":  statusText,
		"message": message,
		"data":    data,
	})
}
