package library

import (
	"github.com/afissard/goblib/server/internal/database"
	"github.com/afissard/goblib/shared/api"
)

func (m *Manager) CreateChapter(chapter database.Chapter) (*api.Chapter, error) {
	created, err := m.db.CreateChapter(chapter)
	if err != nil {
		return nil, err
	}

	apiChapter := m.db.ToApiChapter(created)
	return &apiChapter, nil
}

func (m *Manager) GetChapterByID(id string) (*api.Chapter, error) {
	chapter, err := m.db.GetChapterById(id)
	if err != nil {
		return nil, err
	}

	apiChapter := m.db.ToApiChapter(chapter)
	return &apiChapter, nil
}

func (m *Manager) GetChapterByTitle(title string) (*api.Chapter, error) {
	chapter, err := m.db.GetChapterByTitle(title)
	if err != nil {
		return nil, err
	}

	apiChapter := m.db.ToApiChapter(chapter)
	return &apiChapter, nil
}

func (m *Manager) DeleteChapterById(id string) error {
	return m.db.DeleteChapterById(id)
}

func (m *Manager) UpdateChapter(chapter database.Chapter) (*api.Chapter, error) {
	updated, err := m.db.UpdateChapter(chapter)
	if err != nil {
		return nil, err
	}

	apiChapter := m.db.ToApiChapter(updated)
	return &apiChapter, nil
}

func (m *Manager) ListChapters() (*[]api.Chapter, error) {
	Chapters, err := m.db.ListChapters()
	if err != nil {
		return nil, err
	}

	apiChapters := make([]api.Chapter, len(Chapters))
	for i, b := range Chapters {
		apiChapters[i] = m.db.ToApiChapter(b)
	}
	return &apiChapters, nil
}

func (m *Manager) GetChaptersByBookId(bookId string) (*[]api.Chapter, error) {
	Chapters, err := m.db.GetChaptersByBookId(bookId)
	if err != nil {
		return nil, err
	}

	apiChapters := make([]api.Chapter, len(Chapters))
	for i, b := range Chapters {
		apiChapters[i] = m.db.ToApiChapter(b)
	}
	return &apiChapters, nil
}
