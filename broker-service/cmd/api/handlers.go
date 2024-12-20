package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type RequestPayload struct {
	Action string      `json:"action"`
	User   UserPayload `json:"user,omitempty"`
}

type UserPayload struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := JsonResponse{
		Error:   false,
		Message: "Broker service is running",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}

// HandleSubmission is the main point of entry into the broker. It accepts a JSON
// payload and performs an action based on the value of "action" in that JSON.
func (app *Config) HandleSubmission(w http.ResponseWriter, r *http.Request) {
	var requestPayload RequestPayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	switch requestPayload.Action {
	case "user":
		app.handleUserRequest(w, requestPayload.User)
	}
}

func (app *Config) handleUserRequest(w http.ResponseWriter, user UserPayload) {
	jsonData, _ := json.MarshalIndent(user, "", "\t")

	userServiceURL := "http://user-service/user"

	request, err := http.NewRequest("POST", userServiceURL, bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err, http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusCreated {
		app.errorJSON(w, err)
		return
	}

	// JSON response'dan gelen body'yi işleyelim
	var responsePayload JsonResponse
	err = json.NewDecoder(response.Body).Decode(&responsePayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	// Response'dan gelen user ID'yi almak
	userID, ok := responsePayload.Data.(string)
	if !ok {
		app.errorJSON(w, errors.New("invalid response format"), http.StatusInternalServerError)
		return
	}

	var payload JsonResponse
	payload.Error = false
	payload.Message = "User created successfully"
	payload.Data = userID

	app.writeJSON(w, http.StatusCreated, payload)
}
