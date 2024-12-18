package main

import (
	"log"
	"net/http"
)

const webPort = ":80"

type Config struct{}

func main() {
	app := Config{}

	log.Printf("Starting the application on port %s\n", webPort)

	// define the http server
	srv := &http.Server{
		Addr:    webPort,
		Handler: app.routes(),
	}

	// start the server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
