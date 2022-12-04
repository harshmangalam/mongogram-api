package auth

import (
	"context"
	"mongogram/database"
	"mongogram/models"
	"mongogram/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SignupBody struct {
	Email    string          `json:"email" validate:"required,email"`
	Phone    string          `json:"phone" validate:"required"`
	Name     string          `json:"name" validate:"required"`
	Username string          `json:"username" validate:"required"`
	Password string          `json:"password" validate:"required"`
	Birthday *utils.DateBody `json:"birthday" validate:"required"`
}

func Signup(c *fiber.Ctx) error {

	signupBody := new(SignupBody)

	if err := c.BodyParser(signupBody); err != nil {
		return utils.InternalServerErrorResponse(c, err)
	}

	// validate user input
	errors := utils.ValidateStruct(signupBody)
	if errors != nil {
		return utils.UnprocessedInputResponse(c, fiber.Map{
			"errors": errors,
		})
	}

	usersColl := database.Mi.Db.Collection(database.UsersCollection)
	user := new(models.User)

	// verify duplicate email
	if err := usersColl.FindOne(context.TODO(), bson.D{{Key: "email", Value: signupBody.Email}}).Decode(user); err != nil {
		if err != mongo.ErrNoDocuments {
			return utils.InternalServerErrorResponse(c, err)
		}
	}

	if user.Email != "" {
		return utils.BadRequestErrorResponse(c, "Email already exists")
	}

	// verify duplicate phone number
	if err := usersColl.FindOne(context.TODO(), bson.D{{Key: "phone", Value: signupBody.Phone}}).Decode(user); err != nil {
		if err != mongo.ErrNoDocuments {
			return utils.InternalServerErrorResponse(c, err)
		}
	}
	if user.Phone != "" {
		return utils.BadRequestErrorResponse(c, "Phone number already exists")
	}

	// verify duplicate username
	if err := usersColl.FindOne(context.TODO(), bson.D{{Key: "username", Value: signupBody.Username}}).Decode(user); err != nil {
		if err != mongo.ErrNoDocuments {
			return utils.InternalServerErrorResponse(c, err)
		}
	}
	if user.Username != "" {
		return utils.BadRequestErrorResponse(c, "Username already exists")
	}

	// verify user age (age>18yr)
	// parse string time
	age, birthTime := utils.GetAge(signupBody.Birthday)
	if age <= 18 {
		return utils.BadRequestErrorResponse(c, "Your age must be greater than 18")
	}

	// hash plain password
	hashPassword, err := utils.HashPassword(signupBody.Password)
	if err != nil {
		return utils.InternalServerErrorResponse(c, err)
	}

	// save data
	doc := bson.D{
		{Key: "email", Value: signupBody.Email},
		{Key: "phone", Value: signupBody.Phone},
		{Key: "name", Value: signupBody.Name},
		{Key: "username", Value: signupBody.Username},
		{Key: "birthday", Value: birthTime},
		{Key: "password", Value: hashPassword},
		{Key: "createdAt", Value: time.Now().UTC()},
		{Key: "updatedAt", Value: time.Now().UTC()},
	}
	insertedUser, err := usersColl.InsertOne(context.TODO(), doc)

	if err != nil {
		return utils.InternalServerErrorResponse(c, err)
	}
	// send Verification code on mobile phone

	// send verification code on email

	return utils.CreatedResponse(c, "Account created", fiber.Map{
		"userId": insertedUser.InsertedID,
	})
}
