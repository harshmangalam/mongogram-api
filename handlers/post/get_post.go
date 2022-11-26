package post

import (
	"mongogram/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetPost(c *fiber.Ctx) error {
	postId, err := primitive.ObjectIDFromHex(c.Params("postId"))

	if err != nil {
		return utils.ReturnError(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	post, err := utils.FindPostById(postId)
	if err == nil && post == nil {
		return utils.ReturnError(c, fiber.StatusNotFound, "Post not found", nil)
	}
	if err != nil {
		return utils.ReturnError(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "et post",
		"data": fiber.Map{
			"post": post,
		},
	})
}
