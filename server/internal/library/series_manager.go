package library

import (
	"github.com/afissard/goblib/server/internal/database"
	"github.com/afissard/goblib/shared/api"
)

func (m *Manager) CreateSeries(series database.Series) (*api.Series, error) {
	created, err := m.db.CreateSeries(series)
	if err != nil {
		return nil, err
	}

	apiSeries := m.db.ToApiSeries(created)
	return &apiSeries, nil
}

func (m *Manager) GetSeriesByID(id string) (*api.Series, error) {
	series, err := m.db.GetSeriesById(id)
	if err != nil {
		return nil, err
	}

	apiSeries := m.db.ToApiSeries(series)
	return &apiSeries, nil
}

func (m *Manager) GetSeriesByName(name string) (*api.Series, error) {
	series, err := m.db.GetSeriesByName(name)
	if err != nil {
		return nil, err
	}

	apiSeries := m.db.ToApiSeries(series)
	return &apiSeries, nil
}

func (m *Manager) DeleteSeriesById(id string) error {
	return m.db.DeleteSeriesById(id)
}

func (m *Manager) UpdateSeries(series database.Series) (*api.Series, error) {
	updated, err := m.db.UpdateSeries(series)
	if err != nil {
		return nil, err
	}

	apiSeries := m.db.ToApiSeries(updated)
	return &apiSeries, nil
}

func (m *Manager) ListSeries() (*[]api.Series, error) {
	series, err := m.db.ListSeries()
	if err != nil {
		return nil, err
	}

	apiSeries := make([]api.Series, len(series))
	for i, s := range series {
		apiSeries[i] = m.db.ToApiSeries(s)
	}
	return &apiSeries, nil
}
