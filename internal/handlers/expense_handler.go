package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"go-tracker/internal/dto"
	"go-tracker/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ExpenseHandler struct {
	service *service.ExpenseService
}

func NewExpenseHandler(s *service.ExpenseService) *ExpenseHandler {
	return &ExpenseHandler{
		service: s,
	}
}

// gin.Context is essentially going to be an object that will be providing us with all the information(http method, url, headers, body etc.) and tools(function such as ShouldBindJSON, JSON etc.) for handling one HTTP request, it gets created for EVERY request that hits
func (h *ExpenseHandler) CreateExpense(c *gin.Context) {
	//Set up variable for request.body
	var dto dto.CreateExpenseDTO

	// This is SUPER IMPORTANT. Gin is going to expose a set of methods, one of them being ShouldBindJSON(&targetObject). This will try to parse the body of the request into the actual expense variable we just created, we then use this specific syntax check for error and then just return a 400 Bad Request error
	if err := c.ShouldBindJSON(&dto); err != nil { // shorthand error check syn
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	expense, err := h.service.CreateExpense(dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not create expense",
		})

		return
	}

	c.JSON(http.StatusCreated, expense)
}

func (h *ExpenseHandler) UpdateExpense(c *gin.Context) {
	var dto dto.UpdateExpenseDTO

	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})

		return
	}

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	expense, err := h.service.UpdateExpense(id, dto)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "expense not found",
			})

			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error updating expense",
		})

		return
	}

	c.JSON(http.StatusCreated, expense)
}

func (h *ExpenseHandler) FindAll(c *gin.Context) {
	expenses, err := h.service.FindAllExpenses()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "there has been an error getting expenses",
		})

		return
	}

	c.JSON(http.StatusOK, expenses)
}

func (h *ExpenseHandler) FindById(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})

		return
	}

	expense, err := h.service.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "expense not found",
			})

			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "there has been an error getting the expense",
		})

		return
	}

	c.JSON(http.StatusOK, expense)
}
