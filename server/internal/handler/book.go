package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/afissard/goblib/server/internal/database"
	"github.com/afissard/goblib/server/internal/library"
	shared "github.com/afissard/goblib/shared/logger"
)

type BookHandler struct {
	Manager *library.Manager
	Logger  *shared.Logger
}

func NewBookHandler(manager *library.Manager, logger *shared.Logger) *BookHandler {
	return &BookHandler{
		Manager: manager,
		Logger:  logger,
	}
}

func (h *BookHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		h.Logger.LogMessage("Bookhandler.Get : missing book id", shared.LogLevelError)
		http.Error(w, "missing book id", http.StatusBadRequest)
		return
	}

	book, err := h.Manager.Get(id)
	if err != nil {
		if errors.Is(err, database.ErrBookNotFound) {
			h.Logger.LogMessage("Bookhandler.Get : not found", shared.LogLevelError)
			http.Error(w, "book not found", http.StatusNotFound)
			return
		}

		h.Logger.LogMessage("Bookhandler.Get : internal server error while searching for book", shared.LogLevelError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(book); err != nil {
		h.Logger.LogMessage("Bookhandler.Get : internal server error while encoding response", shared.LogLevelError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	h.Logger.LogMessage("Bookhandler.Get : wrote response", shared.LogLevelInfo)
}

func (h *BookHandler) List(w http.ResponseWriter, r *http.Request) {
	books, err := h.Manager.List()
	if err != nil {
		h.Logger.LogMessage("Bookhandler.List : error while searching for books", shared.LogLevelError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(books); err != nil {
		h.Logger.LogMessage("Bookhandler.List : error while encoding response", shared.LogLevelError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	h.Logger.LogMessage("Bookhandler.List : wrote response", shared.LogLevelInfo)
}

func (h *BookHandler) Create(w http.ResponseWriter, r *http.Request) {
	h.Manager.Create(w, r)
	h.Logger.LogMessage("Bookhandler.Create : wrote response", shared.LogLevelInfo)
}

func (h *BookHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "missing book id", http.StatusBadRequest)
		return
	}
	_, err := h.Manager.Get(id)
	if err != nil {
		if errors.Is(err, database.ErrBookNotFound) {
			http.Error(w, "book not found", http.StatusNotFound)
			return
		}
		h.Logger.LogMessage("Bookhandler.Update : internal server error while searching for book", shared.LogLevelError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.Manager.Update(w, r)
	h.Logger.LogMessage("Bookhandler.Update : wrote response", shared.LogLevelInfo)
}

func (h *BookHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		h.Logger.LogMessage("Bookhandler.Delete : missing book id", shared.LogLevelError)
		http.Error(w, "missing book id", http.StatusBadRequest)
		return
	}
	err := h.Manager.Delete(id)
	if err != nil {
		if errors.Is(err, database.ErrBookNotFound) {
			h.Logger.LogMessage("Bookhandler.Delete : book not found", shared.LogLevelError)
			http.Error(w, "book not found", http.StatusNotFound)
		} else {
			h.Logger.LogMessage("Bookhandler.Delete : internal server error while deleting book", shared.LogLevelError)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
