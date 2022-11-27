package post

import (
	"mongogram/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetPost(c *fiber.Ctx) error {
	postId, err := primitive.ObjectIDFromHex(c.Params("postId"))

	if err != nil {
		return utils.BadRequestErrorResponse(c, "Invalid post id")
	}

	post, err := utils.FindPostById(postId)
	if err == nil && post == nil {
		return utils.NotFoundErrorResponse(c)
	}
	if err != nil {
		return utils.InternalServerErrorResponse(c, err)
	}

	return utils.OkResponse(c, "Get post details", fiber.Map{
		"post": post,
	})
}
