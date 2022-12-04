package accounts

import (
	"context"
	"mongogram/database"
	"mongogram/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

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
