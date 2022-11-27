package accounts

import (
	"context"
	"mongogram/database"
	"mongogram/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type ChangePasswordBody struct {
	CurrentPassword string `json:"currentPassword" validate:"required"`
	NewPassword     string `json:"newPassword" validate:"required"`
}

func ChangePassword(c *fiber.Ctx) error {
	userId := c.Locals("userId")

	changePassword := new(ChangePasswordBody)

	// parse request body
	if err := c.BodyParser(changePassword); err != nil {
		return utils.InternalServerErrorResponse(c, err)

	}

	// validate user input
	error := utils.ValidateStruct(changePassword)

	if error != nil {
		return utils.UnprocessedInputResponse(c, fiber.Map{"errors": error})
	}

	user, err := utils.FindUserById(userId)

	if err == nil && user == nil {
		return utils.NotFoundErrorResponse(c)
	}
	if err != nil {
		return utils.InternalServerErrorResponse(c, err)
	}

	// match current password
	if match := utils.CheckPasswordHash(changePassword.CurrentPassword, user.Password); !match {
		return utils.BadRequestErrorResponse(c, "Current password is incorrect")
	}

	// hash new passord
	hash, err := utils.HashPassword(changePassword.NewPassword)

	if err != nil {
		return utils.InternalServerErrorResponse(c, err)

	}

	// save new password
	usersColl := database.Mi.Db.Collection(database.UsersCollection)
	update := bson.D{
		{"$set", bson.D{
			{"password", hash},
		},
		},
	}
	_, err = usersColl.UpdateByID(context.TODO(), userId, update)
	if err != nil {
		return utils.InternalServerErrorResponse(c, err)
	}

	return utils.OkResponse(c, "Password changed", nil)
}
