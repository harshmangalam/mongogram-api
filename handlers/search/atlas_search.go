package search

import (
	"context"
	"mongogram/database"
	"mongogram/models"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func AtlasSearch(c *fiber.Ctx) error {
	queryText := c.Query("q")
	userId := c.Locals("userId")

	if strings.Trim(queryText, " ") == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Search query must be required",
			"data":    nil,
		})
	}

	usersColl := database.Mi.Db.Collection(database.UsersCollection)
	searchColl := database.Mi.Db.Collection(database.SearchCollection)

	// save text as a recent search

	searchDoc := bson.D{
		{Key: "text", Value: queryText},
		{Key: "userId", Value: userId},
		{Key: "searchedAt", Value: time.Now().UTC()},
	}
	_, err := searchColl.InsertOne(context.TODO(), searchDoc)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
	}

	// create atlas search pipeline
	pipeline := mongo.Pipeline{
		bson.D{
			{"$search", bson.D{
				{"index", database.AtlasSearchIndex},
				{"text", bson.D{
					{"query", queryText},
					{"path", bson.D{
						{"wildcard", "*"},
					}},
				}},
			}},
		},
	}
	cursor, err := usersColl.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
	}
	var users []models.User

	if err = cursor.All(context.TODO(), &users); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "fetch search results",
		"data": fiber.Map{
			"users": users,
		},
	})
}
