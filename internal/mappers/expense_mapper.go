package mappers

import (
	"go-tracker/internal/db"
	dto "go-tracker/internal/dto/expense"
)

func ToExpenseResponse(dbExpense db.Expense) dto.ExpenseResponse {
	return dto.ExpenseResponse{
		ID:        int(dbExpense.ID),
		Title:     dbExpense.Title,
		Amount:    dbExpense.Amount,
		CreatedAt: dbExpense.CreatedAt.Time,
		UpdatedAt: dbExpense.UpdatedAt.Time,
	}
}
