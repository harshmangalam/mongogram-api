package post

import (
	"context"
	"mongogram/database"
	"mongogram/models"
	"mongogram/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetPosts(c *fiber.Ctx) error {
	var posts []models.Post

	postsColl := database.Mi.Db.Collection(database.PostsCollection)

	cursor, err := postsColl.Find(context.TODO(), bson.M{})

	if err != nil {
		return utils.ReturnError(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	if err := cursor.All(context.TODO(), &posts); err != nil {
		return utils.ReturnError(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "get posts",
		"data": fiber.Map{
			"posts": posts,
		},
	})
}
