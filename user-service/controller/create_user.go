package controller

import (
	"encoding/json"
	"net/http"

	"github.com/social-media/user-service/dto"
)

func CreateUser(writer http.ResponseWriter, request *http.Request) {
	var userRequest dto.UserRequest

	// Decode the JSON request body into the reqPayload struct
	if err := json.NewDecoder(request.Body).Decode(&userRequest); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
}
