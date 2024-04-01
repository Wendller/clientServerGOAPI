package database

import (
	"context"
	"database/sql"
	"log"

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

func InsertCotation(ctx context.Context, db *sql.DB, pid string) error {
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
