package api

import (
	"net/http"

	"github.com/afissard/goblib/server/internal/app"
	"github.com/afissard/goblib/server/internal/handler"
	"github.com/afissard/goblib/shared/logger"
)

func NewRouter(a *app.App) http.Handler {
	mux := http.NewServeMux()
	handler := handlers.NewHandler(a.Library, a.Logger)

	// mux.Handle("/", http.FileServer(http.Dir("internal/web")))
	mux.Handle("/", http.FileServer(http.Dir("internal/web/dist"))) // serve the React page

	// Hello
	mux.HandleFunc("GET /hello", handlers.Hello)

	// Author
	mux.HandleFunc("GET /api/authors", handler.ListAuthors)
	mux.HandleFunc("GET /api/authors/{id}", handler.GetAuthorById)
	mux.HandleFunc("GET /api/authors/{name}", handler.GetAuthorByName)
	mux.HandleFunc("POST /api/authors", handler.CreateAuthor)
	mux.HandleFunc("PUT /api/authors/{id}", handler.UpdateAuthor)
	mux.HandleFunc("DELETE /api/authors/{id}", handler.DeleteAuthor)

	// Language
	mux.HandleFunc("GET /api/languages", handler.ListLanguages)
	mux.HandleFunc("GET /api/languages/{id}", handler.GetLanguageById)
	mux.HandleFunc("GET /api/languages/{name}", handler.GetLanguageByName)
	mux.HandleFunc("POST /api/languages", handler.CreateLanguage)
	mux.HandleFunc("PUT /api/languages/{id}", handler.UpdateLanguage)
	mux.HandleFunc("DELETE /api/languages/{id}", handler.DeleteLanguage)

	// Series
	mux.HandleFunc("GET /api/series", handler.ListSeries)
	mux.HandleFunc("GET /api/series/{id}", handler.GetSeriesById)
	mux.HandleFunc("GET /api/series/{name}", handler.GetSeriesByName)
	mux.HandleFunc("POST /api/series", handler.CreateSeries)
	mux.HandleFunc("PUT /api/series/{id}", handler.UpdateSeries)
	mux.HandleFunc("DELETE /api/series/{id}", handler.DeleteSeriesById)

	// Book
	mux.HandleFunc("GET /api/books", handler.ListBooks)
	mux.HandleFunc("GET /api/books/{id}", handler.GetBookById)
	mux.HandleFunc("GET /api/books/{title}", handler.GetBookByTitle)
	mux.HandleFunc("POST /api/books", handler.CreateBook)
	mux.HandleFunc("PUT /api/books/{id}", handler.UpdateBook)
	mux.HandleFunc("DELETE /api/books/{id}", handler.DeleteBookById)

	// Chapters
	mux.HandleFunc("GET /api/chapters", handler.ListChapters)
	mux.HandleFunc("GET /api/chapters/{id}", handler.GetChapterById)
	mux.HandleFunc("GET /api/chapters/{title}", handler.GetChapterByTitle)
	mux.HandleFunc("POST /api/chapters", handler.CreateChapter)
	mux.HandleFunc("PUT /api/chapters/{id}", handler.UpdateChapter)
	mux.HandleFunc("DELETE /api/chapters/{id}", handler.DeleteChapterById)

	a.Logger.LogMessage("Routes are set up", logger.LogLevelDebug)
	return mux
}
