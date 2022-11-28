package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MediaMetadata struct {
	UserId primitive.ObjectID `bson:"userId" json:"userId"`
}
type Media struct {
	Id         primitive.ObjectID `bson:"_id" json:"_id"`
	ChunkSize  int                `bson:"chunkSize" json:"chunkSize"`
	Filename   string             `bson:"filename" json:"filename"`
	Length     int                `bson:"length" json:"length"`
	Metadata   *MediaMetadata     `bson:"metadata" json:"metadata"`
	UploadDate time.Time          `bson:"uploadDate" json:"uploadDate"`
}
