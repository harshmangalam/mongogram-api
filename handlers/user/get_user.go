package user

import (
	"mongogram/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUser(c *fiber.Ctx) error {

	userId, _ := primitive.ObjectIDFromHex(c.Params("userId"))

	user, err := utils.FindUserById(userId)
	if user == nil && err == nil {
		return utils.NotFoundErrorResponse(c)
	}
	if err != nil {
		return utils.InternalServerErrorResponse(c, err)
	}
	return utils.OkResponse(c, "Get user details", fiber.Map{
		"user": user,
	})
}
