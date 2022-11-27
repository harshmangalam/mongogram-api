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
		return utils.InternalServerErrorResponse(c, err)
	}

	if err := cursor.All(context.TODO(), &posts); err != nil {
		return utils.InternalServerErrorResponse(c, err)
	}

	return utils.OkResponse(c, "Get post lists", fiber.Map{
		"posts": posts,
	})
}
