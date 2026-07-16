package api

import (
	"net/http"

	"github.com/afissard/goblib/server/internal/app"
	"github.com/afissard/goblib/server/internal/handler"
	"github.com/afissard/goblib/shared/logger"
)

func NewRouter(a *app.App) http.Handler {
	mux := http.NewServeMux()
	bookHandler := handlers.NewBookHandler(a.Library, a.Logger)

	// mux.Handle("/", http.FileServer(http.Dir("internal/web")))
	mux.Handle("/", http.FileServer(http.Dir("internal/web/dist"))) // serve the React page

	// Hello
	mux.HandleFunc("GET /hello", handlers.Hello)

	// Book
	mux.HandleFunc("GET /api/books", bookHandler.List)
	mux.HandleFunc("GET /api/books/{id}", bookHandler.Get)
	mux.HandleFunc("POST /api/books", bookHandler.Create)
	mux.HandleFunc("PUT /api/books/{id}", bookHandler.Update)
	mux.HandleFunc("DELETE /api/books/{id}", bookHandler.Delete)

	a.Logger.LogMessage("Routes are set up", logger.LogLevelDebug)
	return mux
}
