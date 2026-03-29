package models

import "time"

type Expense struct {
	ID        uint    `gorm: "primaryKey"`
	Amount    float64 `gorm: "not null"`
	Title     string  `gorm: "not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
