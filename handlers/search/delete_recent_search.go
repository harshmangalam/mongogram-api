package search

import (
	"context"
	"mongogram/database"
	"mongogram/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteRecentSearch(c *fiber.Ctx) error {
	userId := c.Locals("userId")
	searchId, err := primitive.ObjectIDFromHex(c.Params("searchId"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Search id is not valid",
			"data":    nil,
		})
	}

	searchColl := database.Mi.Db.Collection(database.SearchCollection)
	filter := bson.D{
		{"_id", searchId},
	}
	// only authorized user can delete search
	search := new(models.Search)
	if err := searchColl.FindOne(context.TODO(), filter).Decode(search); err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  "error",
				"message": "Search not found",
				"data":    nil,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
	}

	if search.UserId != userId {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":  "error",
			"message": "Not allowed to delete other users resource",
			"data":    nil,
		})
	}

	_, err = searchColl.DeleteOne(context.TODO(), filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  "error",
				"message": "Search not found",
				"data":    nil,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Deleted",
		"data":    nil,
	})
}

func DeleteRecentSearchs(c *fiber.Ctx) error {
	userId := c.Locals("userId")

	searchColl := database.Mi.Db.Collection(database.SearchCollection)
	filter := bson.D{
		{"userId", userId},
	}
	// only authorized user can delete search

	_, err := searchColl.DeleteMany(context.TODO(), filter)
	if err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Deleted",
		"data":    nil,
	})
}
