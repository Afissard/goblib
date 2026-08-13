package library

import "github.com/afissard/goblib/server/internal/database"

func (m *Manager) CreateAuthor(author database.Author) (*database.Author, error) {
	created, err := m.db.CreateAuthor(author)
	if err != nil {
		return nil, err
	}

	return &created, nil
}

func (m *Manager) GetAuthorById(id string) (*database.Author, error) {
	author, err := m.db.GetAuthorById(id)
	if err != nil {
		return nil, err
	}
	return &author, nil
}

func (m *Manager) GetAuthorByName(name string) (*database.Author, error) {
	author, err := m.db.GetAuthorByName(name)
	if err != nil {
		return nil, err
	}
	return &author, nil
}

func (m *Manager) DeleteAuthorById(id string) error {
	return m.db.DeleteAuthorById(id)
}

func (m *Manager) UpdateAuthor(author database.Author) (*database.Author, error) {
	updated, err := m.db.UpdateAuthor(&author)
	if err != nil {
		return nil, err
	}
	return &updated, nil
}

func (m *Manager) ListAuthors() (*[]database.Author, error) {
	authors, err := m.db.ListAuthors()
	if err != nil {
		return nil, err
	}
	return &authors, nil
}
