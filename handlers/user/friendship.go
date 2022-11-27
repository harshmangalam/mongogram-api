package user

import (
	"context"
	"mongogram/database"
	"mongogram/models"
	"mongogram/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Friendship(c *fiber.Ctx) error {
	currentUserId := c.Locals("userId")
	usersColl := database.Mi.Db.Collection(database.UsersCollection)

	// first check the user to whome you want to follow exists in db
	// verify user id
	otherUserId, err := primitive.ObjectIDFromHex(c.Params("userId"))
	if err != nil {
		return utils.BadRequestErrorResponse(c, "Invalid user id")
	}
	otherUser, err := utils.FindUserById(otherUserId)

	if otherUser == nil && err == nil {
		return utils.NotFoundErrorResponse(c)
	}

	if err != nil {
		return utils.InternalServerErrorResponse(c, err)
	}

	// check if you already follow other user
	user := new(models.User)
	followerFilter := bson.M{
		"_id": otherUserId,
		"followers": bson.A{
			currentUserId,
		},
	}

	if err := usersColl.FindOne(context.TODO(), followerFilter).Decode(user); err != nil {
		if err == mongo.ErrNoDocuments {

			// you are not following other user

			// do not follow yourself
			if currentUserId == otherUserId {
				return utils.BadRequestErrorResponse(c, "You are not allowed to follow yourself")
			}

			// follow other user
			updateQuery := bson.M{
				"$push": bson.M{
					"followers": bson.A{
						currentUserId,
					},
				},
			}
			_, err = usersColl.UpdateByID(context.TODO(), otherUserId, updateQuery)
			if err != nil {
				return utils.InternalServerErrorResponse(c, err)
			}
			return utils.OkResponse(c, "Follow", nil)

		} else {
			return utils.InternalServerErrorResponse(c, err)
		}
	}

	// you already follow other user
	// unfollow other user

	updateQuery := bson.M{
		"$pull": bson.M{
			"followers": bson.A{
				currentUserId,
			},
		},
	}
	_, err = usersColl.UpdateByID(context.TODO(), otherUserId, updateQuery)
	if err != nil {
		return utils.InternalServerErrorResponse(c, err)
	}

	return utils.OkResponse(c, "Unfollow", nil)
}
