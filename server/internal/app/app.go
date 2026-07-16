package app

import (
	"github.com/afissard/goblib/server/internal/database"
	"github.com/afissard/goblib/server/internal/library"

	shared "github.com/afissard/goblib/shared/logger"
)

type App struct {
	DB      *database.Database
	Library *library.Manager
	Logger  *shared.Logger
}
