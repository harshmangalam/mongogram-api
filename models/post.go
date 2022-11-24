package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	Id        primitive.ObjectID `bson:"_id" json:"_id"`
	MediaType string             `bson:"mediaType" json:"mediaType"`
	MediaUrl  string             `bson:"mediaUrl" json:"mediaUrl"`
	Content   string             `bson:"content" json:"content"`
	AuthorId  primitive.ObjectID `bson:"authorId" json:"authorId"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`
}
