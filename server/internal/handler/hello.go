package handlers

import (
	"encoding/json"
	"net/http"

	shared "github.com/afissard/goblib/shared/api"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	resp := shared.HelloResponse{
		Message: "Hello from the server!",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
