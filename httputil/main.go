package httputil

import (
	"encoding/json"
	"net/http"
)

// Json encodes a generic http response
func Json(w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

type jsonError struct {
	Error string `json:"error"`
}

func Error(w http.ResponseWriter, message string, code int) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jsonError{message})
}
