package accounts

import (
	"mongogram/utils"

	"github.com/gofiber/fiber/v2"
)

type VerifiOtpBody struct {
	Otp   string `json:"otp" validate:"required"`
	Email string `json:"email" validate:"required"`
}

func VerifyOtp(c *fiber.Ctx) error {
	body := new(VerifiOtpBody)
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

	if user.ResetPassOtp != body.Otp {
		return utils.ReturnError(c, fiber.StatusBadRequest, "Incorrect otp", nil)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Otp verified",
		"data":    nil,
	})

}
