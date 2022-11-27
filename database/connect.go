package database

import (
	"context"
	"log"
	"mongogram/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbName = config.Config("MONGODB_NAME")
var mongoURI = config.Config("MONGODB_URI")

func ConnectMongo() error {
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return err
	}

	db := client.Database(dbName)

	gridFsOpts := options.GridFSBucket().SetName("media")
	bucket, err := gridfs.NewBucket(db, gridFsOpts)
	if err != nil {
		return err
	}

	Mi = MongoInstance{
		Client: client,
		Db:     db,
		Bucket: bucket,
	}

	log.Printf("Mongodb connected (%v)", dbName)

	return nil

}
