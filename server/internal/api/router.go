package api

import (
	"net/http"

	"github.com/afissard/goblib/server/internal/handler"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /hello", handlers.Hello)

	mux.Handle("/", http.FileServer(http.Dir("internal/web")))

	return mux
}
