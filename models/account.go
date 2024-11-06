package models

// Account represents a user account
type Account struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Name    string `json:"name"`
	Balance int64  `json:"balance"`
}
