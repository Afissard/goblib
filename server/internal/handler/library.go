package handlers

import "github.com/afissard/goblib/server/internal/library"

type LibraryHandler struct {
	Manager *library.Manager
}

func NewLibraryHandler(manager *library.Manager) *LibraryHandler {
	return &LibraryHandler{
		Manager: manager,
	}
}
