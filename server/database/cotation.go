package database

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/google/uuid"
)

type Cotation struct {
	ID         string
	Bid        string
	InsertedAt string
}

func NewCotation(bid string) *Cotation {
	return &Cotation{
		ID:         uuid.New().String(),
		Bid:        bid,
		InsertedAt: time.Now().String(),
	}
}

func InsertCotation(db *sql.DB, bid string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	cotation := NewCotation(bid)

	query, err := db.PrepareContext(ctx, "INSERT INTO cotations(id, bid, inserted_at) VALUES($1, $2, $3)")
	if err != nil {
		return err
	}
	defer query.Close()

	_, err = query.ExecContext(ctx, cotation.ID, cotation.Bid, cotation.InsertedAt)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			log.Fatal("cotation raw insertion timeout reached")
			return ctx.Err()
		}

		return err
	}

	log.Println("cotation insertion executed successfully")

	return nil
}
