package repository

import (
	"go-tracker/internal/models"

	"gorm.io/gorm"
)

type ExpenseRepository struct {
	DB *gorm.DB
}

func NewExpenseRepository(db *gorm.DB) *ExpenseRepository {
	return &ExpenseRepository{
		DB: db,
	}
}

// this ought to be changed to use models.ExpenseModel, will require some refactoring and the addition of mapper
func (r *ExpenseRepository) Create(expense *models.Expense) error {
	return r.DB.Create(expense).Error
}

// quite frankly the dumbest fucking syntax I have ever encountered in my life
func (r *ExpenseRepository) UpdateExpense(id int, expense *models.Expense) error {
	result := r.DB.Model(&models.Expense{}).
		Where("id = ?", id).
		Updates(expense)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	err := r.DB.First(expense, id).Error
	return err
}

// what is interesting with GORM is that, it will take the referenced slice(expenses) and populate it with the results of the query(r.DB.Find), we could in theory just call the query without assigning it to anything and return the "expenses" slice we passed to it from the FindAll() repo method
func (r *ExpenseRepository) FindAll() ([]models.Expense, error) {
	var expenses []models.Expense
	err := r.DB.Find(&expenses).Error

	return expenses, err
}

func (r *ExpenseRepository) FindById(id int) (*models.Expense, error) {
	var expense *models.Expense

	//First === TypeORM.FindOne()
	err := r.DB.First(&expense, id).Error

	return expense, err
}
