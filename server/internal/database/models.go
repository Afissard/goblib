package database

type Book struct {
	ID string

	Title   string
	Author  string
	Summary string

	Language  string
	SourceURL string

	RootPath  string
	CoverPath string
}
