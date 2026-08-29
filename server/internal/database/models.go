package database

import "github.com/afissard/goblib/shared/api"

type Series struct {
	ID string

	Title          string
	AuthorID       string
	Summary        string
	SourceURL      string
	CoverImagePath string
}

type Book struct {
	ID       string
	SeriesID string

	Title          string
	AuthorID       string
	Summary        string
	LanguageID     string
	CoverImagePath string
}

type Chapter struct {
	ID     string
	BookID string

	Title   string
	Summary string

	// TODO: Add more fields for chapter, such as content, page count, etc.
}

type Author struct {
	ID   string
	Name string
	// Maybe add more fields later
}

type Language struct {
	ID   string
	Name string
}

func ToApiAuthor(dbAuthor Author) api.Author {
	return api.Author{
		ID:   dbAuthor.ID,
		Name: dbAuthor.Name,
	}
}

func ToApiLanguage(dbLanguage Language) api.Language {
	return api.Language{
		ID:    dbLanguage.ID,
		Title: dbLanguage.Name,
	}
}

func (d *Database) ToApiSeries(dbSeries Series) api.Series {
	dbAuthor, err := d.GetAuthorById(dbSeries.AuthorID)
	if err != nil {
		dbAuthor = Author{ID: "", Name: "Unknown"} // Fallback to a default author
	}

	dbBooks, err := d.GetBooksBySeriesId(dbSeries.ID)
	if err != nil {
		dbBooks = []Book{} // Fallback to an empty list
	}

	return api.Series{
		ID:             dbSeries.ID,
		Title:          dbSeries.Title,
		Author:         ToApiAuthor(dbAuthor),
		Summary:        dbSeries.Summary,
		SourceURL:      dbSeries.SourceURL,
		CoverImagePath: dbSeries.CoverImagePath,
		Books:          d.ToApiBooks(dbBooks),
	}
}

func (d *Database) ToApiBooks(dbBooks []Book) []api.Book {
	apiBooks := make([]api.Book, len(dbBooks))
	for i, dbBook := range dbBooks {
		apiBooks[i] = d.ToApiBook(dbBook)
	}
	return apiBooks
}

func (d *Database) ToApiBook(dbBook Book) api.Book {
	dbAuthor, err := d.GetAuthorById(dbBook.AuthorID)
	if err != nil {
		dbAuthor = Author{ID: "", Name: "Unknown"} // Fallback to a default author
	}

	dbLanguage, err := d.GetLanguageById(dbBook.LanguageID)
	if err != nil {
		dbLanguage = Language{ID: "", Name: "Unknown"} // Fallback to a default language
	}

	dbChapters, err := d.GetChaptersByBookId(dbBook.ID)
	if err != nil {
		dbChapters = []Chapter{} // Fallback to an empty list
	}

	return api.Book{
		ID:             dbBook.ID,
		Title:          dbBook.Title,
		Author:         ToApiAuthor(dbAuthor),
		Summary:        dbBook.Summary,
		Language:       ToApiLanguage(dbLanguage),
		CoverImagePath: dbBook.CoverImagePath,
		Chapters:       d.ToApiChapters(dbChapters),
	}
}

func (d *Database) ToApiChapters(dbChapters []Chapter) []api.Chapter {
	apiChapters := make([]api.Chapter, len(dbChapters))
	for i, dbChapter := range dbChapters {
		apiChapters[i] = d.ToApiChapter(dbChapter)
	}
	return apiChapters
}

func (d *Database) ToApiChapter(dbChapter Chapter) api.Chapter {
	return api.Chapter{
		ID:    dbChapter.ID,
		Title: dbChapter.Title,

		// TODO: More fields to be added
	}
}
