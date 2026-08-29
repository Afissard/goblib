package library

import (
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

/*
func (m *Manager) ImportBook(path string) (*database.Book, error) {
	// TODO
	return nil, nil
}

func (m *Manager) CreateBook(w http.ResponseWriter, r *http.Request) {
	// create a book in the database, based on the request
	newBook := &database.Book{
		Title:      r.URL.Query().Get("title"),
		AuthorID:   r.URL.Query().Get("author"),
		Summary:    r.URL.Query().Get("summary"),
		LanguageID: r.URL.Query().Get("language"),
		// TODO: add later a way to import an image for the cover
	}
	err := m.db.CreateBook(newBook)
	if err != nil {
		m.logger.LogMessage(fmt.Sprintf("LibraryManager.Create : internal server error while creating book: %v", newBook), shared.LogLevelError)
		http.Error(w, "Failed to create book", http.StatusInternalServerError)
		return
	}
	m.logger.LogMessage(fmt.Sprintf("LibraryManager.Create : successfully created book: %v", newBook), shared.LogLevelInfo)
}

func (m *Manager) ListBooks() ([]database.Book, error) {
	return m.db.ListBooks()
}

func (m *Manager) GetBookById(id string) (*database.Book, error) {
	return m.db.GetBookById(id)
}

func (m *Manager) DeleteBookById(id string) error {
	return m.db.DeleteBookById(id)
}

func (m *Manager) UpdateBook(w http.ResponseWriter, r *http.Request) {
	updatedBook := &database.Book{
		ID:         r.URL.Query().Get("id"),
		SeriesID:   r.URL.Query().Get("series"),
		Title:      r.URL.Query().Get("title"),
		AuthorID:   r.URL.Query().Get("author"),
		Summary:    r.URL.Query().Get("summary"),
		LanguageID: r.URL.Query().Get("language"),
		// TODO: add later a way to import an image for the cover
	}
	err := m.db.UpdateBook(updatedBook)
	if err != nil {
		m.logger.LogMessage(fmt.Sprintf("LibraryManager.Update : internal server error while updating book: %v", updatedBook), shared.LogLevelError)
		http.Error(w, "Failed to update book", http.StatusInternalServerError)
		return
	}
	m.logger.LogMessage(fmt.Sprintf("LibraryManager.Update : successfully updated book: %v", updatedBook), shared.LogLevelInfo)
}
*/
