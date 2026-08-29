package library

import (
	"github.com/afissard/goblib/server/internal/database"
	"github.com/afissard/goblib/shared/api"
)

func (m *Manager) CreateLanguage(language database.Language) (*api.Language, error) {
	created, err := m.db.CreateLanguage(language)
	if err != nil {
		return nil, err
	}

	apiLanguage := database.ToApiLanguage(created)
	return &apiLanguage, err
}

func (m *Manager) GetLanguageByName(name string) (*api.Language, error) {
	language, err := m.db.GetLanguageByName(name)
	if err != nil {
		return nil, err
	}

	apiLanguage := database.ToApiLanguage(language)
	return &apiLanguage, nil
}

func (m *Manager) GetLanguageById(id string) (*api.Language, error) {
	language, err := m.db.GetLanguageById(id)
	if err != nil {
		return nil, err
	}

	apiLanguage := database.ToApiLanguage(language)
	return &apiLanguage, nil
}

func (m *Manager) DeleteLanguageById(id string) error {
	return m.db.DeleteLanguageById(id)
}

func (m *Manager) UpdateLanguage(language database.Language) (*api.Language, error) {
	updated, err := m.db.UpdateLanguage(language)
	if err != nil {
		return nil, err
	}

	apiLanguage := database.ToApiLanguage(updated)
	return &apiLanguage, nil
}

func (m *Manager) ListLanguages() (*[]api.Language, error) {
	languages, err := m.db.ListLanguages()
	if err != nil {
		return nil, err
	}

	apiLanguages := make([]api.Language, len(languages))
	for i, language := range languages {
		apiLanguages[i] = database.ToApiLanguage(language)
	}
	return &apiLanguages, nil
}
