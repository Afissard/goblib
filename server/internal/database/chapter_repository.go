package database

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

var ErrChapterNotFound = errors.New("chapter not found")

func (d *Database) CreateChapter(chapter Chapter) (Chapter, error) {
	var query = `
INSERT INTO chapters (
	chapter_id,
	book_id,
	title,
	summary
) VALUES (?, ?, ?, ?)
`
	_, err := d.db.Exec(
		query,
		uuid.NewString(),
		chapter.BookID,
		chapter.Title,
		chapter.Summary,
	)
	if err != nil {
		return Chapter{}, err
	}
	return chapter, nil
}

func (d *Database) GetChapterById(id string) (Chapter, error) {
	var query = `
SELECT chapter_id, book_id, title, summary
FROM chapters
WHERE chapter_id = ?
`
	var chapter Chapter
	err := d.db.QueryRow(query, id).Scan(
		&chapter.ID,
		&chapter.BookID,
		&chapter.Title,
		&chapter.Summary,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Chapter{}, ErrChapterNotFound
		}
		return Chapter{}, err
	}
	return chapter, nil
}

func (d *Database) GetChapterByTitle(title string) (Chapter, error) {
	var query = `
SELECT chapter_id, book_id, title, summary
FROM chapters
WHERE title = ?
`
	var chapter Chapter
	err := d.db.QueryRow(query, title).Scan(
		&chapter.ID,
		&chapter.BookID,
		&chapter.Title,
		&chapter.Summary,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Chapter{}, ErrChapterNotFound
		}
		return Chapter{}, err
	}
	return chapter, nil
}

func (d *Database) GetChaptersByBookId(bookID string) ([]Chapter, error) {
	var query = `
SELECT chapter_id, book_id, title, summary
FROM chapters
WHERE book_id = ?
`
	var chapters []Chapter
	rows, err := d.db.Query(query, bookID)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)
	for rows.Next() {
		var chapter Chapter
		err = rows.Scan(
			&chapter.ID,
			&chapter.BookID,
			&chapter.Title,
			&chapter.Summary)
		if err != nil {
			return nil, err
		}
		chapters = append(chapters, chapter)
	}
	return chapters, nil
}

func (d *Database) ListChapters() ([]Chapter, error) {
	var query = `
SELECT chapter_id, book_id, title, summary
FROM chapters
`
	var chapters []Chapter
	rows, err := d.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)
	for rows.Next() {
		var chapter Chapter
		err = rows.Scan(
			&chapter.ID,
			&chapter.BookID,
			&chapter.Title,
			&chapter.Summary)
		if err != nil {
			return nil, err
		}
		chapters = append(chapters, chapter)
	}
	return chapters, nil
}

func (d *Database) DeleteChapterById(id string) error {
	var query = `
DELETE FROM chapters
WHERE chapter_id = ?
`
	_, err := d.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) UpdateChapter(chapter Chapter) (Chapter, error) {
	var query = `
UPDATE chapters
SET book_id = ?, title = ?, summary = ?
WHERE chapter_id = ?
`
	_, err := d.db.Exec(query, chapter.BookID, chapter.Title, chapter.Summary, chapter.ID)
	if err != nil {
		return Chapter{}, err
	}
	return chapter, nil
}
