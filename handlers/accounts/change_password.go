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
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "error",
			"message": err.Error(),
			"data":    nil,
		})
	}

	user, err := utils.FindUserById(userId)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "error",
			"message": err.Error(),
			"data":    nil,
		})
	}

	if user == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"type":    "error",
			"message": "User not found",
			"data":    nil,
		})
	}

	// match current password

	if match := utils.CheckPasswordHash(changePassword.CurrentPassword, user.Password); !match {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"type":    "error",
			"message": "Current password is incorrect",
			"data":    nil,
		})
	}

	// hash new passord

	hash, err := utils.HashPassword(changePassword.NewPassword)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "error",
			"message": err.Error(),
			"data":    nil,
		})

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
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "error",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"type":    "success",
		"message": "Password changed",
		"data":    nil,
	})
}
