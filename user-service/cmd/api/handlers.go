package main

import (
	"log"
	"net/http"
	"user-service/internal/model"
)

type JSONPayload struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (app *Config) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var requestPayload JSONPayload
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	log.Printf("Request payload: %+v\n", requestPayload)

	// insert data
	user := model.User{
		Name:     requestPayload.Name,
		Email:    requestPayload.Email,
		Password: requestPayload.Password,
	}

	log.Printf("User: %+v\n", user)
	err = app.Models.User.CreateUser(user)
	if err != nil {
		app.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	resp := JsonResponse{
		Error:   false,
		Message: "User created successfully",
	}

	app.writeJSON(w, http.StatusCreated, resp)
}
