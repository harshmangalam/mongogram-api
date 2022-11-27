package accounts

import (
	"mongogram/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type EditAccountBody struct {
	Email    string `json:"email" validate:"required,email"`
	Phone    string `json:"phone" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Bio      string `json:"bio"`
}

func EditAccount(c *fiber.Ctx) error {
	userId := c.Locals("userId")
	editAccBody := new(EditAccountBody)

	// parse request body
	if err := c.BodyParser(editAccBody); err != nil {
		return utils.InternalServerErrorResponse(c, err)
	}

	// validate input body
	errors := utils.ValidateStruct(editAccBody)
	if errors != nil {
		return utils.UnprocessedInputResponse(c, fiber.Map{"errors": errors})

	}

	// update account
	updateDoc := bson.M{
		"$set": bson.M{
			"email":    editAccBody.Email,
			"name":     editAccBody.Name,
			"phone":    editAccBody.Phone,
			"username": editAccBody.Username,
			"bio":      editAccBody.Bio,
		},
	}
	_, err := utils.UpdateUser(userId, updateDoc)
	if err != nil {
		return utils.InternalServerErrorResponse(c, err)
	}

	return utils.OkResponse(c, "Account updated", nil)
}
