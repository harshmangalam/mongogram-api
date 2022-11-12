package auth

import (
	"context"
	"math"
	"mongogram/database"
	"mongogram/models"
	"mongogram/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Birthday struct {
	Day   int `json:"day" validate:"required"`
	Month int `json:"month" validate:"required"`
	Year  int `json:"year" validate:"required"`
}
type SignupBody struct {
	Email    string    `json:"email" validate:"required,email"`
	Phone    string    `json:"phone" validate:"required"`
	Name     string    `json:"name" validate:"required"`
	Username string    `json:"username" validate:"required"`
	Password string    `json:"password" validate:"required"`
	Birthday *Birthday `json:"birthday" validate:"required"`
}

func Signup(c *fiber.Ctx) error {

	signupBody := new(SignupBody)

	if err := c.BodyParser(signupBody); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "error",
			"message": err.Error(),
			"data":    nil,
		})
	}

	// validate user input
	errors := utils.ValidateStruct(signupBody)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid input",
			"data": fiber.Map{
				"errors": errors,
			},
		})

	}

	usersColl := database.Mi.Db.Collection(database.UsersCollection)
	user := new(models.User)

	// verify duplicate email
	if err := usersColl.FindOne(context.TODO(), bson.D{{Key: "email", Value: signupBody.Email}}).Decode(user); err != nil {
		if err != mongo.ErrNoDocuments {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"message": err.Error(),
				"data":    nil,
			})
		}
	}

	if user.Email != "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"data":    nil,
			"message": "Email already exists",
		})
	}

	// verify duplicate phone number

	if err := usersColl.FindOne(context.TODO(), bson.D{{Key: "phone", Value: signupBody.Phone}}).Decode(user); err != nil {
		if err != mongo.ErrNoDocuments {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"message": err.Error(),
				"data":    nil,
			})
		}
	}

	if user.Phone != "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Phone number already exists",
			"status":  "error",
			"data":    nil,
		})
	}

	// verify duplicate username
	if err := usersColl.FindOne(context.TODO(), bson.D{{Key: "username", Value: signupBody.Username}}).Decode(user); err != nil {
		if err != mongo.ErrNoDocuments {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
				"status":  "error",
				"data":    nil,
			})
		}
	}

	if user.Username != "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Username already exists",
			"status":  "error",
			"data":    nil,
		})
	}

	// verify user age (age>18yr)

	// parse string time

	birthTime := time.Date(signupBody.Birthday.Year, time.Month(signupBody.Birthday.Month), signupBody.Birthday.Day, 0, 0, 0, 0, time.UTC)
	// calculate user age
	const SecondsInYear = 3.156e+7
	age := math.Round(time.Since(birthTime).Seconds() / SecondsInYear)

	if age <= 18 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Your age must be greater than 18",
			"status":  "error",
			"data":    nil,
		})
	}

	// hash plain password

	hashPassword, err := utils.HashPassword(signupBody.Password)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
	}

	// save data

	doc := bson.D{
		{Key: "email", Value: signupBody.Email},
		{Key: "phone", Value: signupBody.Phone},
		{Key: "name", Value: signupBody.Name},
		{Key: "username", Value: signupBody.Username},
		{Key: "birthday", Value: birthTime},
		{Key: "password", Value: hashPassword},
		{Key: "createdAt", Value: time.Now()},
		{Key: "updatedAt", Value: time.Now()},
	}
	insertedUser, err := usersColl.InsertOne(context.TODO(), doc)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"status":  "error",
			"data":    nil,
		})
	}
	// send Verification code on mobile phone

	// send verification code on email

	return c.JSON(fiber.Map{
		"message": "Account created successfully",
		"status":  "success",
		"data": fiber.Map{
			"userId": insertedUser.InsertedID,
		},
	})
}
