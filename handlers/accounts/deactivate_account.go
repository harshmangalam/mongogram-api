package accounts

import (
	"context"
	"mongogram/database"
	"mongogram/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DeactivateTempBody struct {
	From *utils.DateBody `json:"from" validate:"required"`
	To   *utils.DateBody `json:"to" validate:"required"`
}

func DeactivateAccountPermanent(c *fiber.Ctx) error {
	userId := c.Locals("userId")

	usersColl := database.Mi.Db.Collection(database.UsersCollection)

	var result bson.M

	if err := usersColl.FindOneAndDelete(context.TODO(), bson.M{"_id": userId}).Decode(&result); err != nil {

		if err == mongo.ErrNoDocuments {
			return utils.NotFoundErrorResponse(c)
		}
		return utils.InternalServerErrorResponse(c, err)
	}

	return utils.OkResponse(c, "Account deleted permanently", nil)

}

func DeactivateAccountTemporary(c *fiber.Ctx) error {
	userId := c.Locals("userId")

	body := new(DeactivateTempBody)

	// parse request body
	if err := c.BodyParser(body); err != nil {
		return utils.InternalServerErrorResponse(c, err)
	}

	usersColl := database.Mi.Db.Collection(database.UsersCollection)

	var result bson.M

	if err := usersColl.FindOneAndUpdate(context.TODO(), bson.M{"_id": userId}, bson.M{
		"$set": bson.M{
			"deactivate": bson.M{
				"from": utils.GetDate(body.From),
				"to":   utils.GetDate(body.From),
			},
		},
	}).Decode(&result); err != nil {

		if err == mongo.ErrNoDocuments {
			return utils.NotFoundErrorResponse(c)
		}
		return utils.InternalServerErrorResponse(c, err)
	}

	return utils.OkResponse(c, "Account deactivate temporarity", nil)

}
