package handlers

import (
	"encoding/json"
	"net/http"
)

// Routes sets the routes for the web service.
func Routes() {
	http.HandleFunc("/sendjson", SendJSON)
}

// SendJSON returns a simple Json document
func SendJSON(w http.ResponseWriter, r *http.Request) {
	u := struct {
		Name  string
		Email string
	}{
		Name:  "Meng",
		Email: "meng@gmail.com",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(&u)
}
