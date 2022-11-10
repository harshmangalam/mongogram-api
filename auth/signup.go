package auth

import (
	"mongogram/utils"

	"github.com/gofiber/fiber/v2"
)

func Signup(c *fiber.Ctx) error {

	users := utils.Mi.Db.Collection("users")

	return c.JSON(fiber.Map{
		"message": "Signup",
	})
}
