package database

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	Amount   int
	Type     string
	Balance  int
	WalletID uint
	UserID   uint
}

/*

{
  "amount": "1000",
  "type": "inflow/outflow",
  "balance": "300",
  "wallet_id": "1",
  "created_at": "",
  "updated_at": "",
  "deleted_at": ""
}

*/
