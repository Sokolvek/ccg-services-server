package storage

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

// a
func InitDB() {
	clientOpts := options.Client().ApplyURI(
		"mongodb://localhost:27017/")
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	MongoClient = client
}

func GetClient() *mongo.Client {
	return MongoClient
}
