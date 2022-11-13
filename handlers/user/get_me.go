package user

import "github.com/gofiber/fiber/v2"

func GetMe(c *fiber.Ctx) error {

	userId := c.Locals("userId")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "get current user",
		"data": fiber.Map{
			"user": userId,
		},
	})
}
