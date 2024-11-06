package handlers

import (
	"net/http"
	"task-golang-db/models"

	"github.com/gin-gonic/gin"
)

// In-memory data store for accounts
var accounts = []models.Account{
	{ID: 1, Name: "John Doe", Balance: 1000},
}

// TopUpAccount increases the account balance
func TopUpAccount(c *gin.Context) {
	var request struct {
		AccountID uint  `json:"account_id"`
		Amount    int64 `json:"amount"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, account := range accounts {
		if account.ID == request.AccountID {
			accounts[i].Balance += request.Amount
			c.JSON(http.StatusOK, gin.H{"message": "Top up successful", "balance": accounts[i].Balance})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
}

// GetAccountBalance returns the current balance of an account
func GetAccountBalance(c *gin.Context) {
	// Assume the account ID is provided via authentication
	accountID := 1 // Example ID

	for _, account := range accounts {
		if account.ID == uint(accountID) {
			c.JSON(http.StatusOK, gin.H{"balance": account.Balance})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
}

// Implement other handlers, such as account statistics
