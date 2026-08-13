package database

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

var ErrAuthorNotFound = errors.New("author not found")

func (d *Database) CreateAuthor(author Author) (Author, error) {
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
	if err != nil {
		return Author{}, err
	}
	return author, nil
}

func (d *Database) GetAuthorByName(name string) (Author, error) {
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
			return Author{}, ErrAuthorNotFound
		}
		return Author{}, err
	}
	return author, nil
}

func (d *Database) GetAuthorById(id string) (Author, error) {
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
			return Author{}, ErrAuthorNotFound
		}
		return Author{}, ErrAuthorNotFound
	}
	return author, nil
}

func (d *Database) DeleteAuthorById(id string) error {
	const query = `
DELETE FROM author
WHERE id = ?
`
	_, err := d.db.Exec(query, id)
	return err
}

func (d *Database) UpdateAuthor(author *Author) (Author, error) {
	const query = `
UPDATE author
SET name = ?
WHERE id = ?
`
	_, err := d.db.Exec(query, author.Name, author.ID)
	if err != nil {
		return Author{}, err
	}

	return *author, nil
}

func (d *Database) ListAuthors() ([]Author, error) {
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
	authors := make([]Author, 0)
	for rows.Next() {
		var author Author
		err = rows.Scan(&author.ID, &author.Name)
		if err != nil {
			return nil, err
		}
		authors = append(authors, author)
	}
	return authors, nil
}
