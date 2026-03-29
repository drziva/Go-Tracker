package dto

type CreateExpenseDTO struct {
	Title  string  `json:"title" binding:"required"`
	Amount float64 `json:"amount" binding:"required,gt=0"`
}

type UpdateExpenseDTO struct {
	Title  string  `json:"title" binding:"required"`
	Amount float64 `json:"amount" binding:"required,gt=0"`
}
