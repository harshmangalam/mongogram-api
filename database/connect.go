package database

import (
	"context"
	"log"
	"mongogram/config"

	"go.mongodb.org/mongo-driver/mongo"
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

	Mi = MongoInstance{
		Client: client,
		Db:     db,
	}

	log.Printf("Mongodb connected (%v)", dbName)

	return nil

}
