package accounts

import (
	"context"
	"mongogram/database"
	"mongogram/models"
	"mongogram/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type EditAccountBody struct {
	Email    string `json:"email" validate:"required,email"`
	Phone    string `json:"phone" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Bio      string `json:"bio"`
}

func EditAccount(c *fiber.Ctx) error {
	userId := c.Locals("userId")
	editAccBody := new(EditAccountBody)

	// parse request body
	if err := c.BodyParser(editAccBody); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "error",
			"message": err.Error(),
			"data":    nil,
		})
	}

	// validate input body

	errors := utils.ValidateStruct(editAccBody)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid input",
			"data": fiber.Map{
				"errors": errors,
			},
		})

	}

	usersColl := database.Mi.Db.Collection(database.UsersCollection)

	// update account

	updateDoc := bson.D{
		{"$set", bson.D{
			{"email", editAccBody.Email},
			{"name", editAccBody.Name},
			{"phone", editAccBody.Phone},
			{"username", editAccBody.Username},
			{"bio", editAccBody.Bio},
		},
		},
	}
	user := new(models.User)
	if err := usersColl.FindOneAndUpdate(context.TODO(), bson.D{{"_id", userId}}, updateDoc).Decode(user); err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Account edited",
		"data":    nil,
	})
}
