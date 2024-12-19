package main

import (
	"fmt"
	"log"
	"net/http"
	"user-service/internal/model"
	"user-service/pkg/db"
)

const (
	webPort = "80"
)

type Config struct {
	Models model.Models
}

func main() {
	// Connect to MongoDB and get the client
	client, err := db.ConnectToMongo()
	if err != nil {
		log.Panic("MongoDB connection failed:", err)
	}
	defer db.DisconnectMongo() // Close the connection when the main function exits

	app := Config{
		Models: model.New(client),
	}

	log.Printf("Starting the application on port %s\n", webPort)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Panic("Error while starting the user server:", err)
	}
	// // Get the user
	// retrievedUser, err := user.User.GetUser("67633675597d2a429d5043e4")
	// if err != nil {
	// 	log.Panic("Error while getting user:", err)
	// } else {
	// 	log.Printf("User retrieved: %+v\n", retrievedUser)
	// }
}

/*
	mongodb://admin:password@localhost:64001/logs?authSource=admin&readPreference=primary&appname=MongDB%20Compass&&directConnection=true&ssl=false
*/
