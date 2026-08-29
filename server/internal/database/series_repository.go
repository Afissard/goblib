package database

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

var ErrSeriesNotFound = errors.New("series not found")

func (d *Database) CreateSeries(series Series) (Series, error) {
	var query = `
INSERT INTO series (
    id,
    title,
	author_id,
	summary,
    source_url,
    cover_image_path
)
VALUES (?, ?, ?, ?, ?, ?)
`
	_, err := d.db.Exec(
		query,
		uuid.NewString(),
		series.Title,
		series.AuthorID,
		series.Summary,
		series.SourceURL,
		series.CoverImagePath)

	if err != nil {
		return Series{}, err
	}
	return series, nil
}

func (d *Database) GetSeriesById(id string) (Series, error) {
	var query = `
SELECT
	id,
	title,
	author_id,
	summary,
	source_url,
	cover_image_path
FROM series
WHERE id = ?
`
	var series Series
	err := d.db.QueryRow(query, id).Scan(
		&series.ID,
		&series.Title,
		&series.AuthorID,
		&series.Summary,
		&series.SourceURL,
		&series.CoverImagePath,
	)
	if err != nil {
		return series, ErrSeriesNotFound
	}
	return series, nil
}

func (d *Database) GetSeriesByName(name string) (Series, error) {
	var query = `
SELECT
	id,
	title,
	author_id,
	summary,
	source_url,
	cover_image_path
FROM series
WHERE title = ?
`
	var series Series
	err := d.db.QueryRow(query, name).Scan(
		&series.ID,
		&series.Title,
		&series.AuthorID,
		&series.Summary,
		&series.SourceURL,
		&series.CoverImagePath,
	)
	if err != nil {
		return series, ErrSeriesNotFound
	}
	return series, nil
}

func (d *Database) UpdateSeries(series Series) (Series, error) {
	var query = `
UPDATE series
SET 
    title = ?, 
    author_id = ?, 
    summary = ?, 
    source_url = ?, 
    cover_image_path = ?
WHERE id = ?
`
	_, err := d.db.Exec(
		query,
		series.Title,
		series.AuthorID,
		series.Summary,
		series.SourceURL,
		series.CoverImagePath,
		series.ID)
	if err != nil {
		return Series{}, err
	}
	return series, nil
}

func (d *Database) DeleteSeriesById(id string) error {
	var query = `
DELETE FROM series
WHERE id = ?
`
	_, err := d.db.Exec(query, id)
	return err
}

func (d *Database) ListSeries() ([]Series, error) {
	var query = `
SELECT
	id,
	title,
	author_id,
	summary,
	source_url,
	cover_image_path
FROM series
`
	var series []Series
	rows, err := d.db.Query(query)
	if err != nil {
		return series, err
	}
	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			return
		}
	}(rows)

	for rows.Next() {
		var s Series
		err = rows.Scan(
			&s.ID,
			&s.Title,
			&s.AuthorID,
			&s.Summary,
			&s.SourceURL,
			&s.CoverImagePath,
		)
		if err != nil {
			return series, err
		}
		series = append(series, s)
	}

	return series, nil
}
