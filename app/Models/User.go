package Models

import (
	"errors"
	"gorm.io/gorm"
	"nielscript.com/budgetapp/api/app/Requests"
	"nielscript.com/budgetapp/api/database"
	"time"
)

type User struct {
	ID           string        `json:"-"`
	FirstName    string        `json:"first_name"`
	LastName     string        `json:"last_name"`
	UserName     string        `json:"user_name" gorm:"unique"`
	Email        string        `json:"email" gorm:"unique"`
	PhoneNumber  string        `json:"phone_number" gorm:"unique"`
	Password     string        `json:"-"`
	CreatedAt    time.Time     `json:"created_at"`
	DeletedAt    time.Time     `json:"-"`
	UpdatedAt    time.Time     `json:"-"`
	Wallets      []Wallet      `json:"wallets"`
	Transactions []Transaction `json:"transactions"`
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

func FindUserByEmail(email string) (User, error) {
	var user User
	result := database.DB.Where(&User{Email: email}).First(&user)
	if result.Error != nil {
		if notFound := errors.Is(result.Error, gorm.ErrRecordNotFound); notFound == true {
			return user, errors.New("user not found")
		}
	}

	return user, nil
}

func FindUserByField(fieldName string, fieldValue any) (User, error) {
	var user User
	query := map[string]interface{}{fieldName: fieldValue}
	result := database.DB.Where(query).First(&user)
	if result.Error != nil {
		if notFound := errors.Is(result.Error, gorm.ErrRecordNotFound); notFound == true {
			return user, errors.New("user not found")
		}
	}

	return user, nil
}

// DeleteUser

// UpdateUser

// exists

// wallets

// transactions
