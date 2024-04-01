package database

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/google/uuid"
)

type Cotation struct {
	ID  string
	Pid string
}

func NewCotation(pid string) *Cotation {
	return &Cotation{
		ID:  uuid.New().String(),
		Pid: pid,
	}
}

func InsertCotation(db *sql.DB, pid string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	id := uuid.New().String()

	query, err := db.Prepare("INSERT INTO cotations(id, pid) values($1, $2)")
	if err != nil {
		return err
	}
	defer query.Close()

	_, err = query.ExecContext(ctx, id, pid)
	if err != nil {
		return err
	}

	log.Println("Cotation insertion executed successfully")

	return nil
}
