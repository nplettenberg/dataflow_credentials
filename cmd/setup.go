package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateMongoConnection() (*mongo.Client, error) {

	mongoHost := GetEnv("MONGO_HOST", "localhost")
	mongoPort := GetEnv("MONGO_PORT", "27017")

	mongoUser := GetEnv("MONGO_USERNAME", "root")
	mongoPassword := GetEnv("MONGO_PASSWORD", "root")

	connection := fmt.Sprintf("mongodb://%s:%s@%s:%s", mongoUser, mongoPassword, mongoHost, mongoPort)

	fmt.Printf("Using database connection: %s", connection)

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(connection).SetServerAPIOptions(serverAPI)

	return mongo.Connect(context.TODO(), opts)
}
