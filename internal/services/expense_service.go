package service

import (
	"context"
	"go-tracker/internal/db"
	"go-tracker/internal/dto"
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

// func (s *ExpenseService) FindAllExpenses() ([]models.Expense, error) {
// 	expenses, err := s.repo.FindAll()

// 	return expenses, err
// }

// func (s *ExpenseService) FindById(id int) (*models.Expense, error) {
// 	return s.repo.FindById(id)
// }
