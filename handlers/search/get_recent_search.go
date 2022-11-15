package search

import (
	"context"
	"mongogram/database"
	"mongogram/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetRecentSearch(c *fiber.Ctx) error {
	userId := c.Locals("userId")

	searchColl := database.Mi.Db.Collection(database.SearchCollection)

	var searches []models.Search
	filter := bson.D{
		{"userId", userId},
	}
	cursor, err := searchColl.Find(context.TODO(), filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
	}

	if err = cursor.All(context.TODO(), &searches); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "fetch recent searches",
		"data": fiber.Map{
			"searches": searches,
		},
	})
}
