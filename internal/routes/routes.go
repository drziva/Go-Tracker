package routes

import (
	"go-tracker/internal/handlers"
	"go-tracker/internal/repository"
	"go-tracker/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	expenseRepo := repository.NewExpenseRepository(db)

	expenseService := service.NewExpenseService(expenseRepo)

	expenseHandler := handlers.NewExpenseHandler(expenseService)

	api := r.Group("/api")
	api.POST("/expenses", expenseHandler.CreateExpense)
	api.GET("/expenses", expenseHandler.FindAll)
	api.GET("/expenses/:id", expenseHandler.FindById)
	api.PUT("/expenses/:id", expenseHandler.UpdateExpense)

	return r
}
