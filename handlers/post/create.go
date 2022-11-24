package post

import (
	"mongogram/utils"

	"github.com/gofiber/fiber/v2"
)

type CreatePostBody struct {
	MediaType string `json:"mediaType" validate:"require"`
	MediaUrl  string `json:"mediaUrl" validate:"require"`
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

}
