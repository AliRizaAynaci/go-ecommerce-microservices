package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"user-service/internal/model"
	"user-service/pkg/db"
)

func main() {
	// Connect to MongoDB and get the client
	client, err := db.ConnectToMongo()
	if err != nil {
		log.Panic("MongoDB connection failed:", err)
	}
	defer db.DisconnectMongo()

	// Create a new instance of the User struct
	user := model.New(client)

	// Insert a new user
	ctx := context.Background()
	err = user.User.CreateUser(ctx, model.User{
		Name:     "John Doe",
		Email:    "john.doe@example.com",
		Password: "password",
	})
	if err != nil {
		log.Panic("Error while inserting user:", err)
	} else {
		log.Println("User inserted successfully")
	}

	// Get the user
	retrievedUser, err := user.User.GetUser("67633675597d2a429d5043e4")
	if err != nil {
		log.Panic("Error while getting user:", err)
	} else {
		log.Printf("User retrieved: %+v\n", retrievedUser)
	}

	// Wait for application shutdown signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down the application...")
}

/*
	mongodb://admin:password@localhost:64001/logs?authSource=admin&readPreference=primary&appname=MongDB%20Compass&&directConnection=true&ssl=false
*/
