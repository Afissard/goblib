package handlers

import (
	"github.com/afissard/goblib/server/internal/library"
	shared "github.com/afissard/goblib/shared/logger"
)

type Handler struct {
	Manager *library.Manager
	Logger  *shared.Logger
}

func NewHandler(manager *library.Manager, logger *shared.Logger) *Handler {
	return &Handler{
		Manager: manager,
		Logger:  logger,
	}
}
