package database

import "gorm.io/gorm"

type User struct {
	//ID          uint   `json:"id" gorm:"primary_key"`
	gorm.Model
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	UserName     string `json:"userName" gorm:"index:unique"`
	Email        string `json:"email" gorm:"unique"`
	PhoneNumber  string `json:"phoneNumber" gorm:"unique"`
	Password     string `json:"password"`
	Wallets      []Wallet
	Transactions []Transaction
}
