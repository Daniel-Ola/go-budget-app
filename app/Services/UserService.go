package UserServices

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"nielscript.com/budgetapp/api/app/Models"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CheckAuth(email string, password string) (Models.User, error) {
	user, err := Models.FindUserByEmail(email)
	if err != nil {
		return Models.User{}, err
	}

	if checkPass := checkPasswordHash(password, user.Password); checkPass == false {
		return Models.User{}, errors.New("invalid_credentials")
	}

	return user, nil
}
