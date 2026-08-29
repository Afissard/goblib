package database

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

var ErrBookNotFound = errors.New("book not found")

func (d *Database) CreateBook(book Book) (Book, error) {
	const query = `
INSERT INTO books(
	book_id,
	series_id,
	title,
	author_id,
	summary,
	language_id,
	cover_img_path
)
VALUES (?, ?, ?, ?, ?, ?, ?)
`
	_, err := d.db.Exec(
		query,
		uuid.NewString(),
		book.SeriesID,
		book.Title,
		book.AuthorID,
		book.Summary,
		book.LanguageID,
		book.CoverImagePath,
	)
	if err != nil {
		return Book{}, err
	}
	return book, nil
}

func (d *Database) GetBookById(id string) (Book, error) {

	const query = `
SELECT
    book_id,
	series_id,
	title,
	author_id,
	summary,
	language_id,
	cover_img_path
FROM books
WHERE book_id = ?
`

	var book Book

	err := d.db.QueryRow(query, id).Scan(
		&book.ID,
		&book.SeriesID,
		&book.Title,
		&book.AuthorID,
		&book.Summary,
		&book.LanguageID,
		&book.CoverImagePath,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return Book{}, ErrBookNotFound
	}

	if err != nil {
		return Book{}, err
	}

	return book, nil
}

// GetBookByTitle :
// WARNING: assume that each book title is unique !
func (d *Database) GetBookByTitle(title string) (Book, error) {
	const query = `
SELECT
	book_id,
	series_id,
	title,
	author_id,
	summary,
	language_id,
	cover_img_path
FROM books
WHERE title = ?
`
	var book Book

	err := d.db.QueryRow(query, title).Scan(
		&book.ID,
		&book.SeriesID,
		&book.Title,
		&book.AuthorID,
		&book.Summary,
		&book.LanguageID,
		&book.CoverImagePath,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return Book{}, ErrBookNotFound
	}

	if err != nil {
		return Book{}, err
	}

	return book, nil
}

func (d *Database) GetBooksBySeriesId(seriesId string) ([]Book, error) {
	const query = `
SELECT 
	book_id,
	series_id,
	title,
	author_id,
	summary,
	language_id,
	cover_img_path
FROM books
WHERE series_id = ?
`
	var books []Book

	rows, err := d.db.Query(query, seriesId)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			return
		}
	}(rows)
	for rows.Next() {
		var book Book
		err = rows.Scan(
			&book.ID,
			&book.SeriesID,
			&book.Title,
			&book.AuthorID,
			&book.Summary,
			&book.LanguageID,
			&book.CoverImagePath)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (d *Database) ListBooks() ([]Book, error) {

	const query = `
SELECT
    book_id,
	series_id,
	title,
	author_id,
	summary,
	language_id,
	cover_img_path
FROM books
ORDER BY title
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

	var books []Book

	for rows.Next() {

		var book Book

		err = rows.Scan(
			&book.ID,
			&book.SeriesID,
			&book.Title,
			&book.AuthorID,
			&book.Summary,
			&book.LanguageID,
			&book.CoverImagePath,
		)

		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, rows.Err()
}

func (d *Database) DeleteBookById(id string) error {

	result, err := d.db.Exec(
		`DELETE FROM books WHERE book_id = ?`,
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

func (d *Database) UpdateBook(book Book) (Book, error) {

	result, err := d.db.Exec(
		`
UPDATE books
SET
	series_id = ?,
	title = ?,
	author_id = ?,
	summary = ?,
	language_id = ?,
	cover_img_path = ?
WHERE book_id = ?
`,
		&book.SeriesID,
		&book.Title,
		&book.AuthorID,
		&book.Summary,
		&book.LanguageID,
		&book.CoverImagePath,
		&book.ID,
	)

	if err != nil {
		return Book{}, err
	}

	n, err := result.RowsAffected()
	if err != nil {
		return Book{}, err
	}

	if n == 0 {
		return Book{}, ErrBookNotFound
	}

	return book, nil
}
