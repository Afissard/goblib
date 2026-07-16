package main

import (
	"log"
	"net/http"

	"github.com/afissard/goblib/server/internal/api"
	"github.com/afissard/goblib/server/internal/app"
	"github.com/afissard/goblib/server/internal/database"
	"github.com/afissard/goblib/server/internal/library"
	"github.com/afissard/goblib/shared/logger"
)

func main() {
	// Start logger
	var err error = nil
	logSrv, err := logger.InitLogger(logger.LogLevelDebug, true)
	if err != nil {
		log.Fatal(err)
	}
	defer logSrv.Close()
	go logSrv.Run()

	logSrv.LogMessage("Starting server...", logger.LogLevelDebug)

	// Load data
	err = library.CreateStorageDir()

	db, err := database.Open("data/library.db")
	if err != nil {
		logSrv.LogMessage("Fatal error occurred", logger.LogLevelError)
		log.Fatal(err)
	}
	defer func(db *database.Database) {
		err := db.Close()
		if err != nil {
			logSrv.LogMessage("Error closing database: "+err.Error(), logger.LogLevelError)
		}
	}(db)

	a := &app.App{ // TODO: maybe consider making it a global variable
		DB:      db,
		Library: library.NewManager(db, logSrv),
		Logger:  logSrv,
	}

	// Start router
	router := api.NewRouter(a)

	logSrv.LogMessage("Listening on :8080", logger.LogLevelInfo)
	log.Fatal(http.ListenAndServe(":8080", router))
}
