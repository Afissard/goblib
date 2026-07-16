package library

import (
	"log"
	"os"
)

// CreateStorageDir : Create the data folder, the data/bookx folder and the data/library.db file if they don't exist
func CreateStorageDir() error {
	if err := os.MkdirAll("data/books", 0755); err != nil {
		return err
	}

	if _, err := os.Stat("data/library.db"); os.IsNotExist(err) {
		file, err := os.Create("data/library.db")
		if err != nil {
			log.Fatal(err)
		}
		err = file.Close()
		if err != nil {
			return err
		}
	}

	return nil
}
