package main

import (
	"go-tracker/internal/db"
	"go-tracker/internal/routes"
)

func main() {
	database := db.Connect()

	r := routes.SetupRouter(database)

	r.Run()
}
