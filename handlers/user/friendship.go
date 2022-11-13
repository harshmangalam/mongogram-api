package user

import (
	"context"
	"mongogram/database"
	"mongogram/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Friendship(c *fiber.Ctx) error {
	currentUserId := c.Locals("userId")
	// first check the user to whome you want to follow exists in db
	otherUserId, err := primitive.ObjectIDFromHex(c.Params("userId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid user id",
			"data":    nil,
		})
	}

	usersColl := database.Mi.Db.Collection(database.UsersCollection)

	otherUser := new(models.User)
	if err := usersColl.FindOne(context.TODO(), bson.D{{Key: "_id", Value: otherUserId}}).Decode(otherUser); err != nil {

		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  "error",
				"message": "User to whome you want to follow does not exists",
				"data":    nil,
			})
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"message": err.Error(),
				"data":    nil,
			})
		}
	}

	// check if you already follow other user

	user := new(models.User)
	followerFilter := bson.M{
		"followers": bson.M{
			"$in": bson.A{
				currentUserId,
			},
		},
	}
	if err := usersColl.FindOne(context.TODO(), followerFilter).Decode(user); err != nil {
		if err == mongo.ErrNoDocuments {

			// you are not following other user
			// follow other user

			updateQuery := bson.M{
				"$push": bson.M{
					"followers": bson.A{
						currentUserId,
					},
				},
			}
			updatedResult, err := usersColl.UpdateByID(context.TODO(), otherUserId, updateQuery)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"status":  "error",
					"message": err.Error(),
					"data":    nil,
				})
			}
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"status":  "success",
				"message": "Follow",
				"data":    updatedResult,
			})

		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"message": err.Error(),
				"data":    nil,
			})
		}
	}

	// you are not following other user
	// follow other user

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Unfollow",
		"data":    nil,
	})
}
