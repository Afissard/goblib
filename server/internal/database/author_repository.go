package database

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

var ErrAuthorNotFound = errors.New("author not found")

func (d *Database) CreateAuthor(author Author) error {
	const query = `
INSERT INTO author(
	id,
    name
)
values(?, ?)
`
	_, err := d.db.Exec(
		query,
		uuid.NewString(),
		author.Name)
	return err
}

func (d *Database) GetAuthorByName(name string) (*Author, error) {
	const query = `
SELECT
	id,
 	name
FROM author
WHERE name = ?
`
	var author Author

	err := d.db.QueryRow(query, name).Scan(
		&author.ID,
		&author.Name,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrAuthorNotFound
		}
		return nil, err
	}
	return &author, nil
}

func (d *Database) GetAuthorByID(id int) (*Author, error) {
	const query = `
SELECT
	id,
	name
FROM author
WHERE id = ?
`
	var author Author
	err := d.db.QueryRow(query, id).Scan(
		&author.ID,
		&author.Name,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrAuthorNotFound
		}
	}
	return &author, nil
}

func (d *Database) DeleteAuthorByID(id int) error {
	const query = `
DELETE FROM author
WHERE id = ?
`
	_, err := d.db.Exec(query, id)
	return err
}

func (d *Database) UpdateAuthor(author *Author) error {
	const query = `
UPDATE author
SET name = ?
WHERE id = ?
`
	_, err := d.db.Exec(query, author.Name, author.ID)
	return err
}

func (d *Database) ListAuthors() ([]*Author, error) {
	const query = `
SELECT
	id,
	name
FROM author
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
	authors := make([]*Author, 0)
	for rows.Next() {
		var author Author
		err = rows.Scan(&author.ID, &author.Name)
		if err != nil {
			return nil, err
		}
		authors = append(authors, &author)
	}
	return authors, nil
}

func (d *Database) DeleteAuthor(id string) error {
	const query = `
DELETE FROM author
WHERE id = ?
`
	_, err := d.db.Exec(query, id)
	return err
}
