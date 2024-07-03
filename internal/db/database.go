package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB(filePath string) error {
	var err error
	DB, err = sql.Open("sqlite3", filePath)
	if err != nil {
		return err
	}
	return createTable()
}

func createTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS tickets (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		content TEXT
	);`
	_, err := DB.Exec(query)
	return err
}

func CloseDb() {
	if DB != nil {
		DB.Close()
	}
}
