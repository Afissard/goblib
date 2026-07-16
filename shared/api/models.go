package api

type HelloResponse struct {
	Message string `json:"message"`
}

type Book struct {
	ID string `json:"id"`

	Title   string `json:"title"`
	Author  string `json:"author"`
	Summary string `json:"summary"`

	Language  string `json:"language"`
	SourceURL string `json:"source_url"`

	CoverPath string `json:"cover_path"`
	RootPath  string `json:"root_path"`

	Chapters []Chapter `json:"chapters,omitempty"`
}

type Chapter struct {
	ID string `json:"id"`

	Title    string   `json:"title"`
	PagePath []string `json:"page_path"` // path to where the page's image are stored on the server
}
