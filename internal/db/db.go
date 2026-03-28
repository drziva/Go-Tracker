package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go-tracker/internal/models"
)

// RETURNING db as a pointer ensures we only ever create a single connection pool as well as save space
func Connect() *gorm.DB {
	dsn := "host=localhost user=postgres password=nikola dbname=tracker port = 5432 sslmode=disable" // dsn - Data Source Name

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}) // This RETURNS *db, meaning that we do not have to return &db from this function

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Expense{})

	log.Println("Connected to DB!")

	return db // No need for &db here, it's already a pointer
}
