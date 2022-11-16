package utils

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func DuplicateField(coll *mongo.Collection, field string, value any) (bool, error) {
	data := bson.D{}
	if err := coll.FindOne(context.TODO(), bson.D{{Key: field, Value: value}}).Decode(data); err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
