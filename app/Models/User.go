package Models

import (
	"nielscript.com/budgetapp/api/app/Requests"
	"nielscript.com/budgetapp/api/database"
	"time"
)

type User struct {
	ID           string
	FirstName    string
	LastName     string
	UserName     string `gorm:"unique"`
	Email        string `gorm:"unique"`
	PhoneNumber  string `gorm:"unique"`
	Password     string
	CreatedAt    time.Time
	DeletedAt    time.Time
	UpdatedAt    time.Time
	Wallets      []Wallet
	Transactions []Transaction
}

func CreateNewUser(request Requests.CreateUserRequest) (User, error) {
	user := User{
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		UserName:    request.UserName,
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
		Password:    request.Password,
	}

	err := database.DB.Create(&user)
	if err.Error != nil {
		return user, err.Error
	}
	return user, nil
}

func FindUser(userId uint) (User, error) {
	return User{
		FirstName:   "",
		LastName:    "",
		UserName:    "",
		Email:       "",
		PhoneNumber: "",
		Password:    "",
	}, nil
}

// DeleteUser

// UpdateUser

// exists

// wallets

// transactions
