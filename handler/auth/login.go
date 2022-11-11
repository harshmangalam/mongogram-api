package auth

import "github.com/gofiber/fiber/v2"

func Login(c *fiber.Ctx) error {

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "You have Logged in successfully",
		"data":    fiber.Map{},
	})
}
