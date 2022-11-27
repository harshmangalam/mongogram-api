package accounts

import (
	"mongogram/config"
	"mongogram/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type ResetPasswordBody struct {
	Email string `json:"email" validate:"required"`
}

func ResetPassword(c *fiber.Ctx) error {
	body := new(ResetPasswordBody)
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

	// send otp
	pin, err := utils.GeneratePin(6)
	if err != nil {
		return utils.InternalServerErrorResponse(c, err)
	}

	// save pin in db
	_, err = utils.UpdateUser(user.Id, bson.M{"resetPassOtp": pin})

	if err != nil {
		return utils.InternalServerErrorResponse(c, err)
	}
	message := []byte("To: " + body.Email + "\r\n" +
		"Subject: Reset your password!\r\n" +
		"\r\n" + "Your 6 digit pin code is " + pin +
		"\r\n")
	if err := utils.SendMail([]string{body.Email}, config.Config("SMTP_USERNAME"), message); err != nil {
		return utils.InternalServerErrorResponse(c, err)
	}

	return utils.OkResponse(c, "Send otp in your email", nil)
}
