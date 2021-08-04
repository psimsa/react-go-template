package utils

import (
	"encoding/json"
	"main/models"
	"net/http"
)

func RespondWithJson(w http.ResponseWriter, data models.ApiResponse) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
func RespondWithError(w http.ResponseWriter, status int, data models.ApiResponse) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
