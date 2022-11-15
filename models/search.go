package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Search struct {
	Id         primitive.ObjectID `bson:"_id" json:"_id"`
	Text       string             `bson:"text" json:"text"`
	UserId     primitive.ObjectID `bson:"userId" json:"userId"`
	SearchedAt time.Time          `bson:"serchedAt" json:"serchedAt"`
}
