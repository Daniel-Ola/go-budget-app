package database

import "gorm.io/gorm"

type Wallet struct {
	gorm.Model
	Name         string
	Type         string
	Description  string
	UserID       uint
	Transactions []Transaction
}
