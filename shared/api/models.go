package api

type HelloResponse struct {
	Message string `json:"message"`
}

type Series struct {
	ID string `json:"id"`

	Title          string `json:"title"`
	Author         Author `json:"author"`
	Summary        string `json:"summary"`
	SourceURL      string `json:"source_url"`
	CoverImagePath string `json:"cover_image_path"`
}

type Book struct {
	ID string `json:"id"`

	Title     string    `json:"title"`
	Author    Author    `json:"author"`
	Summary   string    `json:"summary"`
	Language  Language  `json:"language"`
	CoverPath string    `json:"cover_path"`
	Chapters  []Chapter `json:"chapters,omitempty"`
}

type Chapter struct {
	ID string `json:"id"`

	Title    string   `json:"title"`
	PagePath []string `json:"page_path"` // path to where the page's image are stored on the server
}

type Author struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Language struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}
