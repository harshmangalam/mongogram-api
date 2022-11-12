package auth

import (
	"context"
	"mongogram/database"
	"mongogram/models"
	"mongogram/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

	user := new(models.User)
	usersColl := database.Mi.Db.Collection(database.UsersCollection)

	// verify user email/username/phone

	if err := usersColl.FindOne(context.TODO(), bson.D{{Key: "email", Value: loginBody.Username}, {Key: "username", Value: loginBody.Username}, {Key: "phone", Value: loginBody.Username}}).Decode(user); err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":  "error",
				"message": "Invalid credentials",
				"data":    errors,
			})
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"message": err.Error(),
				"data": fiber.Map{
					"fields": loginBody,
				},
			})
		}
	}
	// match user password

	// update user active status

	// create jwt token

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "You have Logged in successfully",
		"data":    fiber.Map{},
	})
}
