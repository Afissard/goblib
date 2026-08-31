package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/afissard/goblib/server/internal/database"
	shared "github.com/afissard/goblib/shared/logger"
)

func httpRequestToSeries(r *http.Request) database.Series {
	return database.Series{
		ID:             r.FormValue("id"),
		Title:          r.FormValue("title"),
		AuthorID:       r.FormValue("author_id"),
		Summary:        r.FormValue("summary"),
		SourceURL:      r.FormValue("source_url"),
		CoverImagePath: r.FormValue("cover_image_path"),
	}
}

func (h *Handler) GetSeriesById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		h.Logger.LogMessage("SeriesHandler.GetSeriesById : missing series id", shared.LogLevelError)
		http.Error(w, "missing series id", http.StatusBadRequest)
		return
	}

	series, err := h.Manager.GetSeriesByID(id)
	if err != nil {
		if errors.Is(err, database.ErrSeriesNotFound) {
			h.Logger.LogMessage("SeriesHandler.GetSeries : Not found", shared.LogLevelError)
			http.Error(w, "series not found", http.StatusNotFound)
			return
		}
		h.Logger.LogMessage("SeriesHandler.GetSeries : failed to get series", shared.LogLevelError)
		http.Error(w, "failed to get series", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(series); err != nil {
		h.Logger.LogMessage("SeriesHandler.GetSeries : failed to encode series", shared.LogLevelError)
		http.Error(w, "failed to encode series", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetSeriesByName(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	if name == "" {
		h.Logger.LogMessage("SeriesHandler.GetSeriesByName : missing series name", shared.LogLevelError)
		http.Error(w, "missing series name", http.StatusBadRequest)
		return
	}

	series, err := h.Manager.GetSeriesByName(name)
	if err != nil {
		if errors.Is(err, database.ErrSeriesNotFound) {
			h.Logger.LogMessage("SeriesHandler.GetSeriesByName : Not found", shared.LogLevelError)
			http.Error(w, "series not found", http.StatusNotFound)
			return
		}
		h.Logger.LogMessage("SeriesHandler.GetSeriesByName : failed to get series", shared.LogLevelError)
		http.Error(w, "failed to get series", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(series); err != nil {
		h.Logger.LogMessage("SeriesHandler.GetSeriesByName : failed to encode series", shared.LogLevelError)
		http.Error(w, "failed to encode series", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) ListSeries(w http.ResponseWriter, r *http.Request) {
	seriesList, err := h.Manager.ListSeries()
	if err != nil {
		h.Logger.LogMessage("SeriesHandler.ListSeries : failed to list series", shared.LogLevelError)
		http.Error(w, "failed to list series", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(seriesList); err != nil {
		h.Logger.LogMessage("SeriesHandler.ListSeries : failed to encode series list", shared.LogLevelError)
		http.Error(w, "failed to encode series list", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) CreateSeries(w http.ResponseWriter, r *http.Request) {
	newSeries := httpRequestToSeries(r)
	series, err := h.Manager.CreateSeries(newSeries)
	if err != nil {
		h.Logger.LogMessage("SeriesHandler.CreateSeries : failed to create series", shared.LogLevelError)
		http.Error(w, "failed to create series", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(series); err != nil {
		h.Logger.LogMessage("SeriesHandler.CreateSeries : failed to encode series", shared.LogLevelError)
		http.Error(w, "failed to encode series", http.StatusInternalServerError)
		return
	}
	h.Logger.LogMessage("SeriesHandler.CreateSeries : successfully created series", shared.LogLevelInfo)
}

func (h *Handler) UpdateSeries(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		h.Logger.LogMessage("SeriesHandler.UpdateSeries : missing series id", shared.LogLevelError)
		http.Error(w, "missing series id", http.StatusBadRequest)
		return
	}

	updatedSeries := httpRequestToSeries(r)
	series, err := h.Manager.UpdateSeries(updatedSeries)
	if err != nil {
		h.Logger.LogMessage("SeriesHandler.UpdateSeries : failed to update series", shared.LogLevelError)
		http.Error(w, "failed to update series", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(series); err != nil {
		h.Logger.LogMessage("SeriesHandler.UpdateSeries : failed to encode series", shared.LogLevelError)
		http.Error(w, "failed to encode series", http.StatusInternalServerError)
		return
	}
	h.Logger.LogMessage("SeriesHandler.UpdateSeries : successfully updated series", shared.LogLevelInfo)
}

func (h *Handler) DeleteSeriesById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		h.Logger.LogMessage("SeriesHandler.DeleteSeriesById : missing series id", shared.LogLevelError)
		http.Error(w, "missing series id", http.StatusBadRequest)
		return
	}

	err := h.Manager.DeleteSeriesById(id)
	if err != nil {
		h.Logger.LogMessage("SeriesHandler.DeleteSeriesById : failed to delete series", shared.LogLevelError)
		http.Error(w, "failed to delete series", http.StatusInternalServerError)
		return
	}
	h.Logger.LogMessage("SeriesHandler.DeleteSeries : successfully deleted series", shared.LogLevelInfo)
}
