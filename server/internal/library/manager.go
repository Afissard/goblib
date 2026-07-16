package library

import (
	"net/http"

	"github.com/afissard/goblib/server/internal/database"
	shared "github.com/afissard/goblib/shared/logger"
)

type Manager struct {
	db     *database.Database
	logger *shared.Logger
}

func NewManager(db *database.Database, logger *shared.Logger) *Manager {
	return &Manager{
		db:     db,
		logger: logger,
	}
}

func (m *Manager) Import(path string) (*database.Book, error) {
	// TODO
	return nil, nil
}

func (m *Manager) Create(w http.ResponseWriter, r *http.Request) {
	// create a book in the database, based on the request
	err := m.db.CreateBook(database.Book{
		Title:     r.URL.Query().Get("title"),
		Author:    r.URL.Query().Get("author"),
		Summary:   r.URL.Query().Get("summary"),
		Language:  r.URL.Query().Get("language"),
		SourceURL: r.URL.Query().Get("source_url"),
		// TODO: add later a way to import an image for the cover
	})
	if err != nil {
		m.logger.LogMessage("LibraryManager.Create : internal server error while creating book", shared.LogLevelError)
		http.Error(w, "Failed to create book", http.StatusInternalServerError)
		return
	}
	m.logger.LogMessage("LibraryManager.Create : successfully created book", shared.LogLevelInfo)
}

func (m *Manager) List() ([]database.Book, error) {
	return m.db.ListBooks()
}

func (m *Manager) Get(id string) (*database.Book, error) {
	return m.db.GetBook(id)
}

func (m *Manager) Delete(id string) error {
	return m.db.DeleteBook(id)
}

func (m *Manager) Update(w http.ResponseWriter, r *http.Request) {

	err := m.db.UpdateBook(database.Book{
		ID:        r.URL.Query().Get("id"),
		Title:     r.URL.Query().Get("title"),
		Author:    r.URL.Query().Get("author"),
		Summary:   r.URL.Query().Get("summary"),
		Language:  r.URL.Query().Get("language"),
		SourceURL: r.URL.Query().Get("source_url"),
		// TODO: add later a way to import an image for the cover
	})
	if err != nil {
		m.logger.LogMessage("LibraryManager.Update : internal server error while updating book", shared.LogLevelError)
		http.Error(w, "Failed to update book", http.StatusInternalServerError)
	}
	m.logger.LogMessage("LibraryManager.Update : successfully updated book", shared.LogLevelInfo)
}

func (m *Manager) DeleteBook(id string) error {
	return m.db.DeleteBook(id)
}
