package library

import (
	"github.com/afissard/goblib/server/internal/database"
	"github.com/afissard/goblib/shared/api"
)

func (m *Manager) CreateBook(book database.Book) (*api.Book, error) {
	created, err := m.db.CreateBook(book)
	if err != nil {
		return nil, err
	}

	apiBook := m.db.ToApiBook(created)
	return &apiBook, nil
}

func (m *Manager) GetBookByID(id string) (*api.Book, error) {
	book, err := m.db.GetBookById(id)
	if err != nil {
		return nil, err
	}

	apiBook := m.db.ToApiBook(book)
	return &apiBook, nil
}

func (m *Manager) GetBookByTitle(title string) (*api.Book, error) {
	book, err := m.db.GetBookByTitle(title)
	if err != nil {
		return nil, err
	}

	apiBook := m.db.ToApiBook(book)
	return &apiBook, nil
}

func (m *Manager) DeleteBookById(id string) error {
	return m.db.DeleteBookById(id)
}

func (m *Manager) UpdateBook(book database.Book) (*api.Book, error) {
	updated, err := m.db.UpdateBook(book)
	if err != nil {
		return nil, err
	}

	apiBook := m.db.ToApiBook(updated)
	return &apiBook, nil
}

func (m *Manager) ListBooks() (*[]api.Book, error) {
	books, err := m.db.ListBooks()
	if err != nil {
		return nil, err
	}

	apiBooks := make([]api.Book, len(books))
	for i, b := range books {
		apiBooks[i] = m.db.ToApiBook(b)
	}
	return &apiBooks, nil
}

func (m *Manager) GetBooksBySeriesId(seriesId string) (*[]api.Book, error) {
	books, err := m.db.GetBooksBySeriesId(seriesId)
	if err != nil {
		return nil, err
	}

	apiBooks := make([]api.Book, len(books))
	for i, b := range books {
		apiBooks[i] = m.db.ToApiBook(b)
	}
	return &apiBooks, nil
}
