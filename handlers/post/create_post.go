package post

import (
	"context"
	"mongogram/database"
	"mongogram/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type CreatePostBody struct {
	MediaType string `json:"mediaType" validate:"required"`
	MediaUrl  string `json:"mediaUrl" validate:"required"`
	Content   string `json:"content"`
}

func CreatePost(c *fiber.Ctx) error {
	body := new(CreatePostBody)
	userId := c.Locals("userId")
	// parse request body
	if err := c.BodyParser(body); err != nil {
		return utils.ReturnError(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	// validate input body

	errors := utils.ValidateStruct(body)
	if errors != nil {
		return utils.ReturnError(c, fiber.StatusBadRequest, "Invalid input", errors)

	}

	postsColl := database.Mi.Db.Collection(database.PostsCollection)

	doc := bson.M{
		"mediaType": body.MediaType,
		"mediaUrl":  body.MediaUrl,
		"content":   body.Content,
		"authorId":  userId,
		"createdAt": time.Now().UTC(),
		"updatedAt": time.Now().UTC(),
	}
	_, err := postsColl.InsertOne(context.TODO(), doc)
	if err != nil {
		return utils.ReturnError(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Post created",
		"data":    nil,
	})
}
