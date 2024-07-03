package db

import (
	"database/sql"

	//_ "github.com/mattn/go-sqlite3"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() error {
	var err error
	connStr := "postgresql://fariz:fariz@localhost:5432/tikecting_master?sslmode=disable"
	DB, err = sql.Open("postgres", connStr)
	//DB, err = sql.Open("sqlite3", filePath)
	if err != nil {
		return err
	}
	return createTable()
}

func createTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS tickets (
		id SERIAL PRIMARY KEY,
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
