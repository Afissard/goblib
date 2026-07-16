package database

import (
	"database/sql"
	"errors"
)

var ErrBookNotFound = errors.New("book not found")

func (d *Database) CreateBook(book Book) error {
	const query = `
INSERT INTO books(
	id,
	title,
	author,
	summary,
	language,
	source_url,
	root_path,
	cover_path
)
VALUES (?, ?, ?, ?, ?, ?, ?, ?)
`

	_, err := d.db.Exec(
		query,
		book.ID,
		book.Title,
		book.Author,
		book.Summary,
		book.Language,
		book.SourceURL,
		book.RootPath,
		book.CoverPath,
	)

	return err
}

func (d *Database) GetBook(id string) (*Book, error) {

	const query = `
SELECT
	id,
	title,
	author,
	summary,
	language,
	source_url,
	root_path,
	cover_path
FROM books
WHERE id = ?
`

	var book Book

	err := d.db.QueryRow(query, id).Scan(
		&book.ID,
		&book.Title,
		&book.Author,
		&book.Summary,
		&book.Language,
		&book.SourceURL,
		&book.RootPath,
		&book.CoverPath,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrBookNotFound
	}

	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (d *Database) ListBooks() ([]Book, error) {

	const query = `
SELECT
	id,
	title,
	author,
	summary,
	language,
	source_url,
	root_path,
	cover_path
FROM books
ORDER BY title
`

	rows, err := d.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var books []Book

	for rows.Next() {

		var book Book

		err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.Summary,
			&book.Language,
			&book.SourceURL,
			&book.RootPath,
			&book.CoverPath,
		)

		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, rows.Err()
}

func (d *Database) DeleteBook(id string) error {

	result, err := d.db.Exec(
		`DELETE FROM books WHERE id = ?`,
		id,
	)

	if err != nil {
		return err
	}

	n, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if n == 0 {
		return ErrBookNotFound
	}

	return nil
}

func (d *Database) UpdateBook(book Book) error {

	result, err := d.db.Exec(
		`
UPDATE books
SET
	title = ?,
	author = ?,
	summary = ?,
	language = ?,
	source_url = ?,
	root_path = ?,
	cover_path = ?
WHERE id = ?
`,
		book.Title,
		book.Author,
		book.Summary,
		book.Language,
		book.SourceURL,
		book.RootPath,
		book.CoverPath,
		book.ID,
	)

	if err != nil {
		return err
	}

	n, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if n == 0 {
		return ErrBookNotFound
	}

	return nil
}
