package service

import (
	"context"
	"go-tracker/internal/db"
	dto "go-tracker/internal/dto/expense"
)

type ExpenseService struct {
	queries *db.Queries
}

func NewExpenseService(queries *db.Queries) *ExpenseService {
	return &ExpenseService{
		queries: queries,
	}
}

func (s *ExpenseService) CreateExpense(dto dto.CreateExpenseDTO) (db.Expense, error) {
	expense, err := s.queries.CreateExpense(
		context.Background(),
		db.CreateExpenseParams{
			Title:  dto.Title,
			Amount: dto.Amount,
		},
	)
	return expense, err
}

func (s *ExpenseService) UpdateExpense(id int, dto dto.UpdateExpenseDTO) (db.Expense, error) {
	expense, err := s.queries.UpdateExpense(
		context.Background(),
		db.UpdateExpenseParams{
			ID:     int32(id),
			Title:  dto.Title,
			Amount: dto.Amount,
		},
	)

	return expense, err
}

func (s *ExpenseService) FindAllExpenses() ([]db.Expense, error) {
	expenses, err := s.queries.GetAllExpenses(context.Background())

	return expenses, err
}

// func (s *ExpenseService) FindById(id int) (*db.Expense, error) {
// 	expense, err := s.queries.GetExpenseByID(context.Background(), )
// }
