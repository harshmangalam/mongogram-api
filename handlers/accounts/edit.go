package accounts

import "github.com/gofiber/fiber/v2"

type EditAccountBody struct {
	Email    string `json:"email" validate:"required,email"`
	Phone    string `json:"phone" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Bio      string `json:"bio"`
}

func EditAccount(c *fiber.Ctx) error {
	editAccBody := new(EditAccountBody)

	// parse request body
	if err := c.BodyParser(editAccBody); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "error",
			"message": err.Error(),
			"data":    nil,
		})
	}

	// validate input body

	// verify user authorization

	// update account

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Account edited",
		"data":    nil,
	})
}
