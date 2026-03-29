package main

import (
	"go-tracker/internal/db"
	"go-tracker/internal/routes"
)

func main() {
	pool := db.NewPostgresPool()

	queries := db.New(pool)

	r := routes.SetupRouter(queries)

	r.Run()
}
