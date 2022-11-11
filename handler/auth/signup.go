package auth

import (
	"context"
	"math"
	"mongogram/database"
	models "mongogram/model"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Birthday struct {
	Day   int `json:"day"`
	Month int `json:"month"`
	Year  int `json:"year"`
}
type SignupBody struct {
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
	Name     string    `json:"name"`
	Username string    `json:"username"`
	Birthday *Birthday `json:"birthday"`
}

func Signup(c *fiber.Ctx) error {

	signupBody := new(SignupBody)

	if err := c.BodyParser(signupBody); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	usersColl := database.Mi.Db.Collection("users")
	user := new(models.User)

	// verify duplicate email
	if err := usersColl.FindOne(context.TODO(), bson.D{{Key: "email", Value: signupBody.Email}}).Decode(user); err != nil {
		if err != mongo.ErrNoDocuments {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
	}

	if user.Email != "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Email already exists",
		})
	}

	// verify duplicate phone number

	if err := usersColl.FindOne(context.TODO(), bson.D{{Key: "phone", Value: signupBody.Phone}}).Decode(user); err != nil {
		if err != mongo.ErrNoDocuments {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
	}

	if user.Phone != "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Phone number already exists",
		})
	}

	// verify duplicate username
	if err := usersColl.FindOne(context.TODO(), bson.D{{Key: "username", Value: signupBody.Username}}).Decode(user); err != nil {
		if err != mongo.ErrNoDocuments {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
	}

	if user.Username != "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Username already exists",
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
		})
	}

	// save data

	doc := bson.D{{Key: "email", Value: signupBody.Email}, {Key: "phone", Value: signupBody.Phone}, {Key: "name", Value: signupBody.Name}, {Key: "username", Value: signupBody.Username}, {Key: "birthday", Value: birthTime}}
	insertedUser, err := usersColl.InsertOne(context.TODO(), doc)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	// send Verification code on mobile phone

	// send verification code on email

	return c.JSON(fiber.Map{
		"message": "Account created successfully",
		"userId":  insertedUser.InsertedID,
	})
}