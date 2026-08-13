package library

import "github.com/afissard/goblib/server/internal/database"

func (m *Manager) CreateLanguage(language database.Language) (*database.Language, error) {
	created, err := m.db.CreateLanguage(language)
	if err != nil {
		return nil, err
	}
	return &created, err
}

func (m *Manager) GetLanguageByName(name string) (database.Language, error) {
	return m.db.GetLanguageByName(name)
}

func (m *Manager) GetLanguageById(id string) (database.Language, error) {
	return m.db.GetLanguageByID(id)
}

func (m *Manager) DeleteLanguageById(id string) error {
	return m.db.DeleteLanguageByID(id)
}

func (m *Manager) UpdateLanguage(language database.Language) (*database.Language, error) {
	updated, err := m.db.UpdateLanguage(&language)
	if err != nil {
		return nil, err
	}
	return &updated, nil
}

func (m *Manager) ListLanguages() (*[]database.Language, error) {
	languages, err := m.db.ListLanguages()
	if err != nil {
		return nil, err
	}
	return &languages, nil
}
