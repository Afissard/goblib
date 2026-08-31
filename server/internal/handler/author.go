package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/afissard/goblib/server/internal/database"
	shared "github.com/afissard/goblib/shared/logger"
)

func httpRequestToAuthor(r *http.Request) database.Author {
	return database.Author{
		ID:   r.FormValue("id"),
		Name: r.FormValue("name"),
	}
}

func (h *Handler) GetAuthorById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		h.Logger.LogMessage("Authorhandler.GetAuthor : missing author id", shared.LogLevelError)
		http.Error(w, "missing author id", http.StatusBadRequest)
		return
	}
	author, err := h.Manager.GetAuthorById(id)
	if err != nil {
		if errors.Is(err, database.ErrAuthorNotFound) {
			h.Logger.LogMessage("Authorhandler.GetAuthor : not found", shared.LogLevelError)
			http.Error(w, "author not found", http.StatusNotFound)
			return
		}
		h.Logger.LogMessage("Authorhandler.GetAuthor : internal server error while searching for author", shared.LogLevelError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(author); err != nil {
		h.Logger.LogMessage("Authorhandler.GetAuthor : internal server error while encoding response", shared.LogLevelError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	h.Logger.LogMessage("Authorhandler.GetAuthor : wrote response", shared.LogLevelInfo)
}

func (h *Handler) GetAuthorByName(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	if name == "" {
		h.Logger.LogMessage("Authorhandler.GetAuthorByName : missing author name", shared.LogLevelError)
		http.Error(w, "missing author name", http.StatusBadRequest)
		return
	}
	author, err := h.Manager.GetAuthorByName(name)
	if err != nil {
		if errors.Is(err, database.ErrAuthorNotFound) {
			h.Logger.LogMessage("Authorhandler.GetAuthorByName : not found", shared.LogLevelError)
			http.Error(w, "author not found", http.StatusNotFound)
			return
		}
		h.Logger.LogMessage("Authorhandler.GetAuthorByName : internal server error while searching for author", shared.LogLevelError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(author); err != nil {
		h.Logger.LogMessage("Authorhandler.GetAuthorByName : internal server error while encoding response", shared.LogLevelError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	h.Logger.LogMessage("Authorhandler.GetAuthorByName : wrote response", shared.LogLevelInfo)
}

func (h *Handler) ListAuthors(w http.ResponseWriter, r *http.Request) {
	authors, err := h.Manager.ListAuthors()
	if err != nil {
		h.Logger.LogMessage("Authorhandler.ListAuthors : error while searching for authors", shared.LogLevelError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(authors); err != nil {
		h.Logger.LogMessage("Authorhandler.ListAuthors : error while encoding response", shared.LogLevelError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	h.Logger.LogMessage("Authorhandler.ListAuthors : wrote response", shared.LogLevelInfo)
}

func (h *Handler) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	newAuthor := httpRequestToAuthor(r)
	author, err := h.Manager.CreateAuthor(newAuthor)
	if err != nil {
		h.Logger.LogMessage("Authorhandler.CreateAuthor : error while creating author", shared.LogLevelError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(author); err != nil {
		h.Logger.LogMessage("Authorhandler.CreateAuthor : error while encoding response", shared.LogLevelError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	h.Logger.LogMessage("Authorhandler.CreateAuthor : wrote response", shared.LogLevelInfo)
}

func (h *Handler) UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		h.Logger.LogMessage("Authorhandler.UpdateAuthor : missing author id", shared.LogLevelError)
		http.Error(w, "missing author id", http.StatusBadRequest)
		return
	}
	_, err := h.Manager.GetAuthorById(id)
	if err != nil {
		if errors.Is(err, database.ErrAuthorNotFound) {
			h.Logger.LogMessage("Authorhandler.UpdateAuthor : not found", shared.LogLevelError)
			http.Error(w, "author not found", http.StatusNotFound)
			return
		}
		h.Logger.LogMessage("Authorhandler.UpdateAuthor : internal server error while searching for author", shared.LogLevelError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	updatedAuthor := httpRequestToAuthor(r)
	updatedAuthor.ID = id
	author, err := h.Manager.UpdateAuthor(updatedAuthor)
	if err != nil {
		h.Logger.LogMessage("Authorhandler.UpdateAuthor : error while updating author", shared.LogLevelError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(author); err != nil {
		h.Logger.LogMessage("Authorhandler.UpdateAuthor : error while encoding response", shared.LogLevelError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	h.Logger.LogMessage("Authorhandler.UpdateAuthor : wrote response", shared.LogLevelInfo)
}

func (h *Handler) DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		h.Logger.LogMessage("Authorhandler.DeleteAuthor : missing author id", shared.LogLevelError)
		http.Error(w, "missing author id", http.StatusBadRequest)
		return
	}
	err := h.Manager.DeleteAuthorById(id)
	if err != nil {
		if errors.Is(err, database.ErrAuthorNotFound) {
			h.Logger.LogMessage("Authorhandler.DeleteAuthor : not found", shared.LogLevelError)
			http.Error(w, "author not found", http.StatusNotFound)
			return
		}
		h.Logger.LogMessage("Authorhandler.DeleteAuthor : internal server error while deleting author", shared.LogLevelError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	h.Logger.LogMessage("Authorhandler.DeleteAuthor : wrote response", shared.LogLevelInfo)
}
