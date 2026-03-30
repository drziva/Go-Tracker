package routes

import (
	"go-tracker/internal/db"
	"go-tracker/internal/handlers"
	service "go-tracker/internal/services"

	"github.com/gin-gonic/gin"
)

func SetupRouter(queries *db.Queries) *gin.Engine {
	r := gin.Default()

	expenseService := service.NewExpenseService(queries)

	expenseHandler := handlers.NewExpenseHandler(expenseService)

	api := r.Group("/api")
	api.POST("/expenses", expenseHandler.CreateExpense)
	api.GET("/expenses", expenseHandler.FindAll)
	// api.GET("/expenses/:id", expenseHandler.FindById)
	api.PUT("/expenses/:id", expenseHandler.UpdateExpense)

	return r
}
