package utils

import (
	"context"
	"mongogram/database"
	"mongogram/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindPostById(postId any) (*models.Post, error) {
	post := new(models.Post)

	postsColl := database.Mi.Db.Collection(database.PostsCollection)

	if err := postsColl.FindOne(context.TODO(), bson.M{"_id": postId}).Decode(post); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return post, nil
}
