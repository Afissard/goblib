package database

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
