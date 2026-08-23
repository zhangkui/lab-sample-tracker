package handler

import (
	"encoding/json"
	"net/http"
)

type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func errorJSON(w http.ResponseWriter, status int, code, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(APIError{code, message})
}
