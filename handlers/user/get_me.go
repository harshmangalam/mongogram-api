package user

import (
	"context"
	"mongogram/database"
	"mongogram/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetMe(c *fiber.Ctx) error {

	userId := c.Locals("userId")

	usersColl := database.Mi.Db.Collection(database.UsersCollection)

	user := new(models.User)
	filter := bson.D{{Key: "_id", Value: userId}}

	if err := usersColl.FindOne(context.TODO(), filter).Decode(user); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "User not found",
			"data":    nil,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "get current user",
		"data": fiber.Map{
			"user": user,
		},
	})
}
