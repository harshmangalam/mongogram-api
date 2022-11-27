package utils

import (
	"context"
	"math"
	"mongogram/database"
	"mongogram/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Birthday struct {
	Day   int `json:"day" validate:"required"`
	Month int `json:"month" validate:"required"`
	Year  int `json:"year" validate:"required"`
}

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

	if err := usersColl.FindOne(context.TODO(), bson.M{"_id": userId}).Decode(user); err != nil {
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

func UpdateUser(id any, update bson.M) (bool, error) {

	usersColl := database.Mi.Db.Collection(database.UsersCollection)

	_, err := usersColl.UpdateByID(context.TODO(), id, update)

	if err != nil {
		return false, err
	}
	return true, nil
}

func GetAge(birthday *Birthday) float64 {
	birthTime := time.Date(birthday.Year, time.Month(birthday.Month), birthday.Day, 0, 0, 0, 0, time.UTC)
	// calculate user age
	const SecondsInYear = 3.156e+7
	return math.Round(time.Since(birthTime).Seconds() / SecondsInYear)

}
