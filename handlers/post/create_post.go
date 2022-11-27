package post

import (
	"context"
	"mongogram/database"
	"mongogram/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreatePostBody struct {
	MediaId string `json:"mediaId" validate:"required"`
	Content string `json:"content"`
}

func CreatePost(c *fiber.Ctx) error {
	body := new(CreatePostBody)
	userId := c.Locals("userId")

	// parse request body
	if err := c.BodyParser(body); err != nil {
		return utils.InternalServerErrorResponse(c, err)
	}

	// validate input body

	errors := utils.ValidateStruct(body)
	if errors != nil {
		return utils.UnprocessedInputResponse(c, fiber.Map{"errors": errors})

	}

	postsColl := database.Mi.Db.Collection(database.PostsCollection)

	mediaId, err := primitive.ObjectIDFromHex(body.MediaId)
	if err != nil {
		return utils.BadRequestErrorResponse(c, "Invalid media id")
	}

	doc := bson.M{
		"mediaId":   mediaId,
		"content":   body.Content,
		"authorId":  userId,
		"createdAt": time.Now().UTC(),
		"updatedAt": time.Now().UTC(),
	}
	createdPost, err := postsColl.InsertOne(context.TODO(), doc)
	if err != nil {
		return utils.InternalServerErrorResponse(c, err)
	}

	return utils.CreatedResponse(c, "Post created", fiber.Map{
		"postId": createdPost.InsertedID,
	})
}
