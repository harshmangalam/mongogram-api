package user

import (
	"mongogram/utils"

	"github.com/gofiber/fiber/v2"
)

func GetMe(c *fiber.Ctx) error {

	userId := c.Locals("userId")

	user, err := utils.FindUserById(userId)
	if user == nil && err == nil {
		return utils.NotFoundErrorResponse(c)
	}
	if err != nil {
		return utils.InternalServerErrorResponse(c, err)
	}
	return utils.OkResponse(c, "Get current user", fiber.Map{
		"user": user,
	})
}
