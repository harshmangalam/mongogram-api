package auth

import (
	"context"
	"mongogram/models"
	"mongogram/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SignupBody struct {
	email    string `json:"email"`
	phone    string `json:"phone"`
	name     string `json:"name"`
	username string `json:"username"`
	birthday string `json:"birthday"`
}

func Signup(c *fiber.Ctx) error {

	signupBody := new(SignupBody)

	if err := c.BodyParser(signupBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Please provide valid inputs",
			"data":    signupBody,
			"error":   err,
		})
	}

	usersColl := utils.Mi.Db.Collection("users")

	// verify duplicate email
	user := new(models.User)
	if err := usersColl.FindOne(context.TODO(), bson.D{{Key: "email", Value: signupBody.email}}).Decode(user); err != nil {
		if err != mongo.ErrNoDocuments {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Error while verifying user email address",
				"data":    signupBody,
				"error":   err,
			})
		}
	}

	// verify duplicate phone number

	if err := usersColl.FindOne(context.TODO(), bson.D{{Key: "phone", Value: signupBody.phone}}).Decode(user); err != nil {
		if err != mongo.ErrNoDocuments {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Error while verifying user phone number",
				"data":    signupBody,
				"error":   err,
			})
		}
	}

	// verify duplicate username
	if err := usersColl.FindOne(context.TODO(), bson.D{{Key: "username", Value: signupBody.username}}).Decode(user); err != nil {
		if err != mongo.ErrNoDocuments {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Error while verifying username",
				"data":    signupBody,
				"error":   err,
			})
		}
	}

	// verify user age (age>18yr)

	// save data

	// send Verification code on mobile phone

	// send verification code on email

	return c.JSON(fiber.Map{
		"message": "Signup",
	})
}
