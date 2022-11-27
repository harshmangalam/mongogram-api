package media

import (
	"context"
	"mongogram/database"
	"mongogram/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetMedia(c *fiber.Ctx) error {

	bucket := database.Mi.Bucket

	filter := bson.M{}
	cursor, err := bucket.Find(filter)

	if err != nil {
		return utils.InternalServerErrorResponse(c, err)
	}

	mediaList := []bson.M{}
	if err = cursor.All(context.TODO(), &mediaList); err != nil {
		return utils.InternalServerErrorResponse(c, err)
	}

	return utils.OkResponse(c, "Get media", fiber.Map{"media": mediaList})
}
