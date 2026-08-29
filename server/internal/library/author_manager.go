package library

import (
	"github.com/afissard/goblib/server/internal/database"
	"github.com/afissard/goblib/shared/api"
)

func (m *Manager) CreateAuthor(author database.Author) (*api.Author, error) {
	created, err := m.db.CreateAuthor(author)
	if err != nil {
		return nil, err
	}

	authorApi := database.ToApiAuthor(created)
	return &authorApi, nil
}

func (m *Manager) GetAuthorById(id string) (*api.Author, error) {
	author, err := m.db.GetAuthorById(id)
	if err != nil {
		return nil, err
	}

	authorApi := database.ToApiAuthor(author)
	return &authorApi, nil
}

func (m *Manager) GetAuthorByName(name string) (*api.Author, error) {
	author, err := m.db.GetAuthorByName(name)
	if err != nil {
		return nil, err
	}
	authorApi := database.ToApiAuthor(author)
	return &authorApi, nil
}

func (m *Manager) DeleteAuthorById(id string) error {
	return m.db.DeleteAuthorById(id)
}

func (m *Manager) UpdateAuthor(author database.Author) (*api.Author, error) {
	updated, err := m.db.UpdateAuthor(author)
	if err != nil {
		return nil, err
	}
	updatedApi := database.ToApiAuthor(updated)
	return &updatedApi, nil
}

func (m *Manager) ListAuthors() (*[]api.Author, error) {
	authors, err := m.db.ListAuthors()
	if err != nil {
		return nil, err
	}

	authorsApi := make([]api.Author, len(authors))
	for i, author := range authors {
		authorsApi[i] = database.ToApiAuthor(author)
	}
	return &authorsApi, nil
}
