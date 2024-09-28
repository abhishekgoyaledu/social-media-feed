package controller

import (
	"encoding/json"
	"net/http"
)

func GetUser(writer http.ResponseWriter, request *http.Request) {

	// Encode the user information as JSON and write it to the response writer
	if err := json.NewEncoder(writer).Encode(response); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
