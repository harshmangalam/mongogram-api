package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	Id        primitive.ObjectID `bson:"_id" json:"_id"`
	MediaId   primitive.ObjectID `bson:"mediaId" json:"mediaId"`
	Content   string             `bson:"content" json:"content"`
	AuthorId  primitive.ObjectID `bson:"authorId" json:"authorId"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`
}
