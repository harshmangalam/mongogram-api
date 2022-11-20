package auth

import (
	"encoding/json"
	"mongogram/config"
	"mongogram/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type AuthBody struct {
	Code string `json:"code" validate:"required"`
}

func LoginWithGithub(c *fiber.Ctx) error {
	// extract code from request body
	authBody := new(AuthBody)

	if err := c.BodyParser(authBody); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
	}

	// validate request body

	errors := utils.ValidateStruct(authBody)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid input",
			"data": fiber.Map{
				"errors": errors,
			},
		})

	}

	// make a post request on github server with code , client id and client secret to get acces token

	clientId := config.Config("GH_CLIENT_ID")
	clientSecret := config.Config("GH_CLIENT_SECRET")
	accessTokenUrl := "https://github.com/login/oauth/access_token" + "?client_id=" + clientId + "&client_secret=" + clientSecret + "&code=" + authBody.Code
	res, err := http.Post(accessTokenUrl, "application/json", nil)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
	}
	defer res.Body.Close()

	var ghRes map[string]any

	if err := json.NewDecoder(res.Body).Decode(&ghRes); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "login successfully",
		"data":    ghRes,
	})
}
