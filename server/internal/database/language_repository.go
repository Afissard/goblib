package database

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

var ErrLanguageNotFound = errors.New("language not found")

func (d *Database) CreateLanguage(language Language) (Language, error) {
	const query = `
INSERT INTO languages(
    id,
	name
)
values(?, ?)
`
	_, err := d.db.Exec(query, uuid.NewString(), language.Name)
	if err != nil {
		return Language{}, err
	}
	return language, nil
}

func (d *Database) GetLanguageByName(name string) (Language, error) {
	const query = `
SELECT
	id,
 	name
FROM languages
WHERE name = ?
`
	var language Language

	err := d.db.QueryRow(query, name).Scan(
		&language.ID,
		&language.Name,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Language{}, ErrLanguageNotFound
		}
		return Language{}, err
	}
	return language, nil
}

func (d *Database) GetLanguageByID(id string) (Language, error) {
	const query = `
SELECT
	id,
	name
FROM languages
WHERE id = ?
`
	var language Language
	err := d.db.QueryRow(query, id).Scan(
		&language.ID,
		&language.Name,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Language{}, ErrLanguageNotFound
		}
		return Language{}, err
	}
	return language, nil
}

func (d *Database) DeleteLanguageByID(id string) error {
	const query = `
DELETE FROM languages
WHERE id = ?
`
	_, err := d.db.Exec(query, id)
	return err
}

func (d *Database) UpdateLanguage(language *Language) (Language, error) {
	const query = `
UPDATE languages
SET name = ?
WHERE id = ?
`
	_, err := d.db.Exec(query, language.Name, language.ID)
	if err != nil {
		return Language{}, err
	}
	return *language, nil
}

func (d *Database) ListLanguages() ([]Language, error) {
	const query = `
SELECT
	id,
	name
FROM languages
ORDER BY name
`
	rows, err := d.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			return
		}
	}(rows)
	languages := make([]Language, 0)
	for rows.Next() {
		var language Language
		err = rows.Scan(&language.ID, &language.Name)
		if err != nil {
			return nil, err
		}
		languages = append(languages, language)
	}
	return languages, nil
}
