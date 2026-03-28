package handlers

import (
	"net/http"

	"go-tracker/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ExpenseHandler struct {
	DB *gorm.DB
}

func NewExpenseHandler(db *gorm.DB) *ExpenseHandler {
	return &ExpenseHandler{
		DB: db,
	}
}

// gin.Context is essentially going to be an object that will be providing us with all the information(http method, url, headers, body etc.) and tools(function such as ShouldBindJSON, JSON etc.) for handling one HTTP request, it gets created for EVERY request that hits
func (h *ExpenseHandler) CreateExpense(c *gin.Context) {
	var expense models.Expense

	// This is SUPER IMPORTANT. Gin is going to expose a set of methods, one of them being ShouldBindJSON(&targetObject). This will try to parse the body of the request into the actual expense variable we just created, we then use this specific syntax check for error and then just return a 400 Bad Request error
	if err := c.ShouldBindJSON(&expense); err != nil { // shorthand error check syn
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	if err := h.DB.Create(&expense).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create expense",
		})

		return
	}

	h.DB.Create(&expense)

	c.JSON(http.StatusCreated, expense)
}
