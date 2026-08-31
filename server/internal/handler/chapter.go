package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/afissard/goblib/server/internal/database"
	shared "github.com/afissard/goblib/shared/logger"
)

func httpRequestToChapter(r *http.Request) database.Chapter {
	return database.Chapter{
		ID:      r.FormValue("id"),
		BookID:  r.FormValue("book_id"),
		Title:   r.FormValue("title"),
		Summary: r.FormValue("summary"),
	}
}

func (h *Handler) GetChapterById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		h.Logger.LogMessage("ChapterHandler.GetChapterById : missing chapter id", shared.LogLevelError)
		http.Error(w, "missing chapter id", http.StatusBadRequest)
		return
	}

	chapter, err := h.Manager.GetChapterByID(id)
	if err != nil {
		if errors.Is(err, database.ErrChapterNotFound) {
			h.Logger.LogMessage("ChapterHandler.GetChapterById : chapter not found", shared.LogLevelError)
			http.Error(w, "chapter not found", http.StatusNotFound)
			return
		}
		h.Logger.LogMessage("ChapterHandler.GetChapterById : failed to get chapter", shared.LogLevelError)
		http.Error(w, "failed to get chapter", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(chapter); err != nil {
		h.Logger.LogMessage("ChapterHandler.GetChapterById : failed to encode chapter", shared.LogLevelError)
		http.Error(w, "failed to encode chapter", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetChapterByTitle(w http.ResponseWriter, r *http.Request) {
	title := r.PathValue("title")
	if title == "" {
		h.Logger.LogMessage("ChapterHandler.GetChapterByTitle : missing chapter title", shared.LogLevelError)
		http.Error(w, "missing chapter title", http.StatusBadRequest)
		return
	}

	chapter, err := h.Manager.GetChapterByTitle(title)
	if err != nil {
		if errors.Is(err, database.ErrChapterNotFound) {
			h.Logger.LogMessage("ChapterHandler.GetChapterByTitle : chapter not found", shared.LogLevelError)
			http.Error(w, "chapter not found", http.StatusNotFound)
			return
		}
		h.Logger.LogMessage("ChapterHandler.GetChapterByTitle : failed to get chapter", shared.LogLevelError)
		http.Error(w, "failed to get chapter", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(chapter); err != nil {
		h.Logger.LogMessage("ChapterHandler.GetChapterByTitle : failed to encode chapter", shared.LogLevelError)
		http.Error(w, "failed to encode chapter", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetChapterByBookId(w http.ResponseWriter, r *http.Request) {
	bookID := r.PathValue("book_id")
	if bookID == "" {
		h.Logger.LogMessage("ChapterHandler.GetChapterByBookId : missing book id", shared.LogLevelError)
		http.Error(w, "missing book id", http.StatusBadRequest)
		return
	}

	chapters, err := h.Manager.GetChaptersByBookId(bookID)
	if err != nil {
		if errors.Is(err, database.ErrChapterNotFound) {
			h.Logger.LogMessage("ChapterHandler.GetChapterByBookId : chapters not found", shared.LogLevelError)
			http.Error(w, "chapters not found", http.StatusNotFound)
			return
		}
		h.Logger.LogMessage("ChapterHandler.GetChapterByBookId : failed to get chapters", shared.LogLevelError)
		http.Error(w, "failed to get chapters", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(chapters); err != nil {
		h.Logger.LogMessage("ChapterHandler.GetChapterByBookId : failed to encode chapters", shared.LogLevelError)
		http.Error(w, "failed to encode chapters", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) ListChapters(w http.ResponseWriter, r *http.Request) {
	chapters, err := h.Manager.ListChapters()
	if err != nil {
		h.Logger.LogMessage("ChapterHandler.ListChapters : failed to list chapters", shared.LogLevelError)
		http.Error(w, "failed to list chapters", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(chapters); err != nil {
		h.Logger.LogMessage("ChapterHandler.ListChapters : failed to encode chapters", shared.LogLevelError)
		http.Error(w, "failed to encode chapters", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) CreateChapter(w http.ResponseWriter, r *http.Request) {
	newChapter := httpRequestToChapter(r)
	if newChapter.BookID == "" || newChapter.Title == "" {
		h.Logger.LogMessage("ChapterHandler.CreateChapter : missing required fields", shared.LogLevelError)
		http.Error(w, "missing required fields", http.StatusBadRequest)
		return
	}

	chapter, err := h.Manager.CreateChapter(newChapter)
	if err != nil {
		h.Logger.LogMessage("ChapterHandler.CreateChapter : failed to create chapter", shared.LogLevelError)
		http.Error(w, "failed to create chapter", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(chapter); err != nil {
		h.Logger.LogMessage("ChapterHandler.CreateChapter : failed to encode chapter", shared.LogLevelError)
		http.Error(w, "failed to encode chapter", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) UpdateChapter(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		h.Logger.LogMessage("ChapterHandler.UpdateChapter : missing chapter id", shared.LogLevelError)
		http.Error(w, "missing chapter id", http.StatusBadRequest)
		return
	}

	updatedChapter := httpRequestToChapter(r)
	if updatedChapter.Title == "" {
		h.Logger.LogMessage("ChapterHandler.UpdateChapter : missing required fields", shared.LogLevelError)
		http.Error(w, "missing required fields", http.StatusBadRequest)
		return
	}

	chapter, err := h.Manager.UpdateChapter(updatedChapter)
	if err != nil {
		h.Logger.LogMessage("ChapterHandler.UpdateChapter : failed to update chapter", shared.LogLevelError)
		http.Error(w, "failed to update chapter", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(chapter); err != nil {
		h.Logger.LogMessage("ChapterHandler.UpdateChapter : failed to encode chapter", shared.LogLevelError)
		http.Error(w, "failed to encode chapter", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) DeleteChapterById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		h.Logger.LogMessage("ChapterHandler.DeleteChapterById : missing chapter id", shared.LogLevelError)
		http.Error(w, "missing chapter id", http.StatusBadRequest)
		return
	}

	err := h.Manager.DeleteChapterById(id)
	if err != nil {
		h.Logger.LogMessage("ChapterHandler.DeleteChapterById : failed to delete chapter", shared.LogLevelError)
		http.Error(w, "failed to delete chapter", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
