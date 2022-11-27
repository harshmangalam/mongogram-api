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
		return utils.InternalServerErrorResponse(c, err)
	}

	// validate input body
	errors := utils.ValidateStruct(body)
	if errors != nil {
		return utils.UnprocessedInputResponse(c, fiber.Map{"errors": errors})

	}

	// verify user account
	user, err := utils.FindUser("email", body.Email)

	if err == nil && user == nil {
		return utils.NotFoundErrorResponse(c)
	}
	if err != nil {
		return utils.InternalServerErrorResponse(c, err)
	}

	if user.ResetPassOtp != body.Otp {
		return utils.BadRequestErrorResponse(c, "Incorrect otp")
	}

	return utils.OkResponse(c, "Otp verified", nil)

}
