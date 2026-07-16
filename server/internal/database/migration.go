package database

func (d *Database) migrate() error {
	const query = `
CREATE TABLE IF NOT EXISTS books (
	id TEXT PRIMARY KEY,

	title TEXT NOT NULL,
	author TEXT,
	summary TEXT,
	language TEXT,
	source_url TEXT,

	root_path TEXT NOT NULL,
	cover_path TEXT,

	created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
`

	_, err := d.db.Exec(query)
	return err
}
