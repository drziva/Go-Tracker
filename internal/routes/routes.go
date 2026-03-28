package routes

import (
	"go-tracker/internal/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	expenseHandler := handlers.NewExpenseHandler(db)

	api := r.Group("/api")
	api.POST("/expenses", expenseHandler.CreateExpense)

	return r
}
