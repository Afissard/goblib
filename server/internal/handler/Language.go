package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/afissard/goblib/server/internal/database"
	shared "github.com/afissard/goblib/shared/logger"
)

func httpRequestToLanguage(r *http.Request) database.Language {
	return database.Language{
		Name: r.FormValue("name"),
	}
}

func (h *Handler) GetLanguageById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		h.Logger.LogMessage("Languagehandler.GetLanguageById : missing language id", shared.LogLevelError)
		http.Error(w, "missing language id", http.StatusBadRequest)
		return
	}
	language, err := h.Manager.GetLanguageById(id)
	if err != nil {
		if err == database.ErrLanguageNotFound {
			h.Logger.LogMessage("Languagehandler.GetLanguageById : not found", shared.LogLevelError)
			http.Error(w, "language not found", http.StatusNotFound)
			return
		}
		h.Logger.LogMessage("Languagehandler.GetLanguageById : internal server error while searching for language", shared.LogLevelError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(language); err != nil {
		h.Logger.LogMessage("Languagehandler.GetLanguageById : internal server error while encoding response", shared.LogLevelError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	h.Logger.LogMessage("Languagehandler.GetLanguageById : wrote response", shared.LogLevelInfo)
}

func (h *Handler) GetLanguageByName(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	if name == "" {
		h.Logger.LogMessage("Languagehandler.GetLanguageByName : missing language name", shared.LogLevelError)
		http.Error(w, "missing language name", http.StatusBadRequest)
		return
	}
	language, err := h.Manager.GetLanguageByName(name)
	if err != nil {
		if err == database.ErrLanguageNotFound {
			h.Logger.LogMessage("Languagehandler.GetLanguageByName : not found", shared.LogLevelError)
			http.Error(w, "language not found", http.StatusNotFound)
			return
		}
		h.Logger.LogMessage("Languagehandler.GetLanguageByName : internal server error while searching for language", shared.LogLevelError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(language); err != nil {
		h.Logger.LogMessage("Languagehandler.GetLanguageByName : internal server error while encoding response", shared.LogLevelError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	h.Logger.LogMessage("Languagehandler.GetLanguageByName : wrote response", shared.LogLevelInfo)
}

func (h *Handler) ListLanguages(w http.ResponseWriter, r *http.Request) {
	languages, err := h.Manager.ListLanguages()
	if err != nil {
		h.Logger.LogMessage("Languagehandler.ListLanguages : error while searching for languages", shared.LogLevelError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(languages); err != nil {
		h.Logger.LogMessage("Languagehandler.ListLanguages : error while encoding response", shared.LogLevelError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	h.Logger.LogMessage("Languagehandler.ListLanguages : wrote response", shared.LogLevelInfo)
}

func (h *Handler) CreateLanguage(w http.ResponseWriter, r *http.Request) {
	language := httpRequestToLanguage(r)
	if language.Name == "" {
		h.Logger.LogMessage("Languagehandler.CreateLanguage : missing language name", shared.LogLevelError)
		http.Error(w, "missing language name", http.StatusBadRequest)
		return
	}
	newLanguage, err := h.Manager.CreateLanguage(language)
	if err != nil {
		h.Logger.LogMessage("Languagehandler.CreateLanguage : error while creating language", shared.LogLevelError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(newLanguage); err != nil {
		h.Logger.LogMessage("Languagehandler.CreateLanguage : error while encoding response", shared.LogLevelError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	h.Logger.LogMessage("Languagehandler.CreateLanguage : wrote response", shared.LogLevelInfo)
}

func (h *Handler) UpdateLanguage(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		h.Logger.LogMessage("Languagehandler.UpdateLanguage : missing language id", shared.LogLevelError)
		http.Error(w, "missing language id", http.StatusBadRequest)
		return
	}
	language := httpRequestToLanguage(r)
	if language.Name == "" {
		h.Logger.LogMessage("Languagehandler.UpdateLanguage : missing language name", shared.LogLevelError)
		http.Error(w, "missing language name", http.StatusBadRequest)
		return
	}
	updatedLanguage, err := h.Manager.UpdateLanguage(language)
	if err != nil {
		if errors.Is(err, database.ErrLanguageNotFound) {
			h.Logger.LogMessage("Languagehandler.UpdateLanguage : not found", shared.LogLevelError)
			http.Error(w, "language not found", http.StatusNotFound)
			return
		}
		h.Logger.LogMessage("Languagehandler.UpdateLanguage : error while updating language", shared.LogLevelError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(updatedLanguage); err != nil {
		h.Logger.LogMessage("Languagehandler.UpdateLanguage : error while encoding response", shared.LogLevelError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.Logger.LogMessage("Languagehandler.UpdateLanguage : wrote response", shared.LogLevelInfo)
}

func (h *Handler) DeleteLanguage(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		h.Logger.LogMessage("Languagehandler.DeleteLanguage : missing language id", shared.LogLevelError)
		http.Error(w, "missing language id", http.StatusBadRequest)
		return
	}
	err := h.Manager.DeleteLanguageById(id)
	if err != nil {
		if errors.Is(err, database.ErrLanguageNotFound) {
			h.Logger.LogMessage("Languagehandler.DeleteLanguage : not found", shared.LogLevelError)
			http.Error(w, "language not found", http.StatusNotFound)
			return
		}
		h.Logger.LogMessage("Languagehandler.DeleteLanguage : error while deleting language", shared.LogLevelError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	h.Logger.LogMessage("Languagehandler.DeleteLanguage : wrote response", shared.LogLevelInfo)
}
