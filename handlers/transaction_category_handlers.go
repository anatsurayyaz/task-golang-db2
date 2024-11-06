package handlers

import (
	"net/http"
	"task-golang-db/models"

	"github.com/gin-gonic/gin"
)

// In-memory data store (for demonstration)
var transactionCategories = []models.TransactionCategory{}

// CreateTransactionCategory handles the creation of a new transaction category
func CreateTransactionCategory(c *gin.Context) {
	var newCategory models.TransactionCategory
	if err := c.ShouldBindJSON(&newCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newCategory.ID = uint(len(transactionCategories) + 1) // Auto-increment ID
	transactionCategories = append(transactionCategories, newCategory)
	c.JSON(http.StatusCreated, newCategory)
}

// ListTransactionCategories returns all transaction categories
func ListTransactionCategories(c *gin.Context) {
	c.JSON(http.StatusOK, transactionCategories)
}

// Implement other handlers: UpdateTransactionCategory and DeleteTransactionCategory
