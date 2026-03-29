package service

import (
	"go-tracker/internal/dto"
	"go-tracker/internal/models"
	"go-tracker/internal/repository"
)

type ExpenseService struct {
	repo *repository.ExpenseRepository
}

func NewExpenseService(repo *repository.ExpenseRepository) *ExpenseService {
	return &ExpenseService{
		repo: repo,
	}
}

func (s *ExpenseService) CreateExpense(dto dto.CreateExpenseDTO) (models.Expense, error) {
	expense := models.Expense{
		Title:  dto.Title,
		Amount: dto.Amount,
	}

	err := s.repo.Create(&expense)

	return expense, err
}

func (s *ExpenseService) UpdateExpense(id int, dto dto.UpdateExpenseDTO) (models.Expense, error) {
	expense := models.Expense{
		Amount: dto.Amount,
		Title:  dto.Title,
	}

	err := s.repo.UpdateExpense(id, &expense)

	return expense, err
}

func (s *ExpenseService) FindAllExpenses() ([]models.Expense, error) {
	expenses, err := s.repo.FindAll()

	return expenses, err
}

func (s *ExpenseService) FindById(id int) (*models.Expense, error) {
	return s.repo.FindById(id)
}
