package auth

import (
	"mongogram/utils"

	"github.com/gofiber/fiber/v2"
)

type SignupBody struct {
	email    string `json:"email"`
	phone    string `json:"phone"`
	name     string `json:"name"`
	username string `json:"username"`
	birthday string `json:"birthday"`
}

func Signup(c *fiber.Ctx) error {

	signupBody := new(SignupBody)

	if err := c.BodyParser(signupBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Please provide valid inputs",
			"data":    signupBody,
		})
	}

	users := utils.Mi.Db.Collection("users")

	// verify duplicate email

	// verify duplicate phone number

	// verify duplicate username

	// verify user age (age>18yr)

	// save data

	// send Verification code on mobile phone

	// send verification code on email

	return c.JSON(fiber.Map{
		"message": "Signup",
	})
}
