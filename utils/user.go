package utils

import (
	"context"
	"mongogram/database"
	"mongogram/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func DuplicateUserField(coll *mongo.Collection, field string, value any) (bool, error) {
	data := bson.D{}
	if err := coll.FindOne(context.TODO(), bson.D{{Key: field, Value: value}}).Decode(data); err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func FindUserById(userId any) (*models.User, error) {
	user := new(models.User)

	usersColl := database.Mi.Db.Collection(database.UsersCollection)

	if err := usersColl.FindOne(context.TODO(), bson.D{{"_id", userId}}).Decode(user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return user, nil
}

func FindUser(field string, value string) (*models.User, error) {
	user := new(models.User)

	usersColl := database.Mi.Db.Collection(database.UsersCollection)

	filter := bson.M{}
	filter[field] = value
	if err := usersColl.FindOne(context.TODO(), filter).Decode(user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return user, nil
}
