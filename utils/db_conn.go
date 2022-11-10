package utils

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var Mi MongoInstance

const dbName = "mongogram"
const mongoURI = "mongodb://localhost:27017/" + dbName

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
