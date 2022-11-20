package utils

import "github.com/gofiber/fiber/v2"

func ReturnError(c *fiber.Ctx, status int, message string, data any) error {
	return c.Status(status).JSON(fiber.Map{
		"status":  "error",
		"message": message,
		"data":    data,
	})
}

func ReturnSuccess(c *fiber.Ctx, status int, message string, data any) error {
	return c.Status(status).JSON(fiber.Map{
		"status":  "success",
		"message": message,
		"data":    data,
	})
}
