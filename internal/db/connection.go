package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresPool() *pgxpool.Pool {
	dsn := "postgres://postgres:nikola@localhost:5432/tracker"

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatal("Failed to connect to DB!", err)
	}

	return pool
}
