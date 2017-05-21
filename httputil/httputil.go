package httputil

import (
	"encoding/json"
	"net/http"
)

// Json returns the response as json
func Json(w http.ResponseWriter, response interface{}, code int) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// jsonError defines the structure of the json message that is returned
type jsonError struct {
	Error string `json:"error"`
}

// Error returns a json error
func Error(w http.ResponseWriter, message string, code int) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jsonError{message})
}
