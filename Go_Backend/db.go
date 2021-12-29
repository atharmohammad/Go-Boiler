package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func getMongoDbConn() (*mongo.Client, error) {
	err := godotenv.Load(".env")
	uri := os.Getenv("MONGODB_URI")
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	return client, err
}

func getDbCollections(dbname string, collectionName string) (*mongo.Collection, error) {
	client, err := getMongoDbConn()

	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database(dbname).Collection(collectionName)
	return collection, nil
}
