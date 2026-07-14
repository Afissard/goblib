package main

import (
	"log"
	"net/http"

	"github.com/afissard/goblib/server/internal/api"
)

func main() {
	router := api.NewRouter()

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
