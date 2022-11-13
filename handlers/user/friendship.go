package user

import "github.com/gofiber/fiber/v2"

func Friendship(c *fiber.Ctx) error {

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "following",
		"data":    nil,
	})
}
