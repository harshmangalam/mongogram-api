package auth

import (
	"mongogram/utils"

	"github.com/gofiber/fiber/v2"
)

type LoginBody struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func Login(c *fiber.Ctx) error {

	loginBody := new(LoginBody)

	if err := c.BodyParser(loginBody); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
			"data": fiber.Map{
				"fields": loginBody,
			},
		})
	}

	// validate users input
	errors := utils.ValidateStruct(loginBody)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid input",
			"data":    errors,
		})

	}
	// verify user email/username/phone

	// match user password

	// update user active status

	// create jwt token

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "You have Logged in successfully",
		"data":    fiber.Map{},
	})
}
