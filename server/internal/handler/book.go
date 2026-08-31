package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/afissard/goblib/server/internal/database"
	shared "github.com/afissard/goblib/shared/logger"
)

func httpRequestToBook(r *http.Request) database.Book {
	return database.Book{
		ID:             r.FormValue("id"),
		SeriesID:       r.FormValue("series_id"),
		Title:          r.FormValue("title"),
		AuthorID:       r.FormValue("author_id"),
		Summary:        r.FormValue("summary"),
		LanguageID:     r.FormValue("language_id"),
		CoverImagePath: r.FormValue("cover_image_path"),
	}
}

func (h *Handler) GetBookById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		h.Logger.LogMessage("BookHandler.GetBookById : missing book id", shared.LogLevelError)
		http.Error(w, "missing book id", http.StatusBadRequest)
		return
	}

	book, err := h.Manager.GetBookByID(id)
	if err != nil {
		if errors.Is(err, database.ErrBookNotFound) {
			h.Logger.LogMessage("BookHandler.GetBookById : book not found", shared.LogLevelError)
			http.Error(w, "book not found", http.StatusNotFound)
			return
		}
		h.Logger.LogMessage("BookHandler.GetBookById : failed to get book", shared.LogLevelError)
		http.Error(w, "failed to get book", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(book); err != nil {
		h.Logger.LogMessage("BookHandler.GetBookById : failed to encode book", shared.LogLevelError)
		http.Error(w, "failed to encode book", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetBookByTitle(w http.ResponseWriter, r *http.Request) {
	title := r.PathValue("title")
	if title == "" {
		h.Logger.LogMessage("BookHandler.GetBookByTitle : missing book title", shared.LogLevelError)
		http.Error(w, "missing book title", http.StatusBadRequest)
		return
	}

	book, err := h.Manager.GetBookByTitle(title)
	if err != nil {
		if errors.Is(err, database.ErrBookNotFound) {
			h.Logger.LogMessage("BookHandler.GetBookByTitle : book not found", shared.LogLevelError)
			http.Error(w, "book not found", http.StatusNotFound)
			return
		}
		h.Logger.LogMessage("BookHandler.GetBookByTitle : failed to get book", shared.LogLevelError)
		http.Error(w, "failed to get book", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(book); err != nil {
		h.Logger.LogMessage("BookHandler.GetBookByTitle : failed to encode book", shared.LogLevelError)
		http.Error(w, "failed to encode book", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) ListBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.Manager.ListBooks()
	if err != nil {
		h.Logger.LogMessage("BookHandler.ListBooks : failed to list books", shared.LogLevelError)
		http.Error(w, "failed to list books", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(books); err != nil {
		h.Logger.LogMessage("BookHandler.ListBooks : failed to encode books", shared.LogLevelError)
		http.Error(w, "failed to encode books", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) CreateBook(w http.ResponseWriter, r *http.Request) {
	newBook := httpRequestToBook(r)
	book, err := h.Manager.CreateBook(newBook)
	if err != nil {
		h.Logger.LogMessage("BookHandler.CreateBook : failed to create book", shared.LogLevelError)
		http.Error(w, "failed to create book", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(book); err != nil {
		h.Logger.LogMessage("BookHandler.CreateBook : failed to encode book", shared.LogLevelError)
		http.Error(w, "failed to encode book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		h.Logger.LogMessage("BookHandler.UpdateBook : missing book id", shared.LogLevelError)
		http.Error(w, "missing book id", http.StatusBadRequest)
		return
	}

	updatedBook := httpRequestToBook(r)
	book, err := h.Manager.UpdateBook(updatedBook)
	if err != nil {
		if errors.Is(err, database.ErrBookNotFound) {
			h.Logger.LogMessage("BookHandler.UpdateBook : book not found", shared.LogLevelError)
			http.Error(w, "book not found", http.StatusNotFound)
			return
		}
		h.Logger.LogMessage("BookHandler.UpdateBook : failed to update book", shared.LogLevelError)
		http.Error(w, "failed to update book", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(book); err != nil {
		h.Logger.LogMessage("BookHandler.UpdateBook : failed to encode book", shared.LogLevelError)
		http.Error(w, "failed to encode book", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) DeleteBookById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		h.Logger.LogMessage("BookHandler.DeleteBookById : missing book id", shared.LogLevelError)
		http.Error(w, "missing book id", http.StatusBadRequest)
		return
	}

	err := h.Manager.DeleteBookById(id)
	if err != nil {
		if errors.Is(err, database.ErrBookNotFound) {
			h.Logger.LogMessage("BookHandler.DeleteBookById : book not found", shared.LogLevelError)
			http.Error(w, "book not found", http.StatusNotFound)
			return
		}
		h.Logger.LogMessage("BookHandler.DeleteBookById : failed to delete book", shared.LogLevelError)
		http.Error(w, "failed to delete book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
