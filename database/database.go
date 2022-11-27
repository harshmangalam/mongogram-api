package database

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
	Bucket *gridfs.Bucket
}

var Mi MongoInstance
