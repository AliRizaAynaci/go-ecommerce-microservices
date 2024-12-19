package main

import (
	"net/http"
	"user-service/internal/model"
)

type JSONPayload struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (app *Config) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	// read the json into var
	var requestPayload struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// insert data
	user := model.User{
		Name:     requestPayload.Name,
		Email:    requestPayload.Email,
		Password: requestPayload.Password,
	}

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
