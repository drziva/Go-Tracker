package dto

import "time"

type CreateExpenseDTO struct {
	Title  string  `json:"title" binding:"required"`
	Amount float64 `json:"amount" binding:"required,gt=0"`
}

type UpdateExpenseDTO struct {
	Title  string  `json:"title" binding:"required"`
	Amount float64 `json:"amount" binding:"required,gt=0"`
}

type ExpenseResponse struct {
	ID        int
	Title     string
	Amount    float64
	CreatedAt time.Time
	UpdatedAt time.Time
}
