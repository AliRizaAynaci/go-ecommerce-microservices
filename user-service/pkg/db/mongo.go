package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	mongoURL = "mongodb://mongo:27017"
	username = "admin"
	password = "password"
)

var client *mongo.Client

const mongoTimeout = 15 * time.Second

func ConnectToMongo() (*mongo.Client, error) {
	// create a connection options
	clientOptions := options.Client().ApplyURI(mongoURL)
	clientOptions.SetAuth(options.Credential{
		Username: username,
		Password: password,
	})

	// Connect
	c, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	client = c

	log.Println("Connected to MongoDB!")
	return client, nil
}

func DisconnectMongo() {
	ctx, cancel := context.WithTimeout(context.Background(), mongoTimeout)
	defer cancel()

	if err := client.Disconnect(ctx); err != nil {
		panic(err)
	}
	log.Println("MongoDB connection closed")
}
