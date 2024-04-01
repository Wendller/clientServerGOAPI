package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func NewCotationsDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "cotations.db")
	if err != nil {
		return nil, fmt.Errorf("🔌 database connection failed: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("⚠️ failed to ping database: %v", err)
	}

	return db, nil
}
