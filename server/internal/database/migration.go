package database

func (d *Database) migrate() error {
	const query = `
CREATE TABLE IF NOT EXISTS Authors (
	id TEXT PRIMARY KEY,
	name TEXT NOT NULL,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS Languages (
	id TEXT PRIMARY KEY,
	name TEXT NOT NULL,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS Series (
    id TEXT PRIMARY KEY,
    
    title TEXT NOT NULL,
    author_id TEXT,
    summary TEXT,
    source_url TEXT,
    cover_image_path TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                                  
	FOREIGN KEY (author_id) REFERENCES Authors(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS books (
    book_id TEXT PRIMARY KEY,
	series_id TEXT,
	
	title TEXT NOT NULL,
	author_id TEXT,
	summary TEXT,
	language_id TEXT,
	cover_img_path TEXT,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	
    FOREIGN KEY (series_id) REFERENCES Series(id) ON DELETE CASCADE,
    FOREIGN KEY (author_id) REFERENCES Authors(id) ON DELETE CASCADE,
    FOREIGN KEY (language_id) REFERENCES Languages(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS chapters (
    chapter_id TEXT PRIMARY KEY,
	book_id TEXT,
	
	title TEXT NOT NULL,
	summary TEXT,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	
	FOREIGN KEY (book_id) REFERENCES books(book_id) ON DELETE CASCADE
);
`
	_, err := d.db.Exec(query)
	return err
}
