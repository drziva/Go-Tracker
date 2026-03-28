package models

import "time"

type Expense struct {
	ID        uint    `gorm: "primaryKey"`
	Title     string  `gorm: "not null"`
	Amount    float64 `grom: not null"`
	CreatedAt time.Time
}
