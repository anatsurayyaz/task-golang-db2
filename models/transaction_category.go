package models

// TransactionCategory represents a category for transactions
type TransactionCategory struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
