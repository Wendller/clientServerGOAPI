package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func NewCotationsDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./cotations.db")
	if err != nil {
		return nil, fmt.Errorf("🔌 database connection failed: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("⚠️ failed to ping database: %v", err)
	}

	err = createTable(db)
	if err != nil {
		return nil, fmt.Errorf("cotations table creation failed: %v", err)
	}

	return db, nil
}

func createTable(db *sql.DB) error {
	query, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS cotations (
			id varchar(255) primary key,
			bid varchar(80),
			inserted_at text
		);
	`)

	if err != nil {
		return err
	}

	_, err = query.Exec()
	if err != nil {
		return err
	}

	return nil
}
