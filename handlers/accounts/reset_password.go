package accounts

import (
	"mongogram/config"
	"mongogram/utils"

	"github.com/gofiber/fiber/v2"
)

type ResetPasswordBody struct {
	Email string `json:"email" validate:"required"`
}

func ResetPassword(c *fiber.Ctx) error {
	body := new(ResetPasswordBody)
	// parse request body
	if err := c.BodyParser(body); err != nil {
		return utils.ReturnError(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	// validate input body

	errors := utils.ValidateStruct(body)
	if errors != nil {
		return utils.ReturnError(c, fiber.StatusBadRequest, "Invalid input", errors)

	}

	// verify user account
	user, err := utils.FindUser("email", body.Email)

	if err == nil && user == nil {
		return utils.ReturnError(c, fiber.StatusNotFound, "User not found", nil)
	}
	if err != nil {
		return utils.ReturnError(c, fiber.StatusInternalServerError, err.Error(), nil)
	}
	// send otp
	pin, err := utils.GeneratePin(6)
	if err != nil {
		return utils.ReturnError(c, fiber.StatusInternalServerError, err.Error(), nil)
	}
	message := []byte("To: " + body.Email + "\r\n" +
		"Subject: Reset your password!\r\n" +
		"\r\n" + "Your 6 digit pin code is " + pin +
		"\r\n")
	if err := utils.SendMail([]string{body.Email}, config.Config("SMTP_USERNAME"), message); err != nil {
		return utils.ReturnError(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Send otp on your email",
		"data":    nil,
	})
}
