package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type CotationDB struct {
	DB *sql.DB
}

func NewCotationsDB() *CotationDB {
	db, err := sql.Open("sqlite3", "./cotations.db")
	if err != nil {
		log.Fatalf("🔌 database connection failed: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("⚠️ failed to ping database: %v", err)
	}

	return &CotationDB{DB: db}
}
