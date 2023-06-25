package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"nielscript.com/budgetapp/api/app/Models"
	"nielscript.com/budgetapp/api/app/Requests"
	UserServices "nielscript.com/budgetapp/api/app/Services"
	"nielscript.com/budgetapp/api/app/Validator"
)

const loginSuccessful, validationError = "User signed in successfully", "Validation Error"
const createUserFailed = "Failed to create user"

func Login(context *gin.Context) {
	gin.BasicAuth(gin.Accounts{})
	var validated Requests.LoginUserRequest
	if err := context.ShouldBindJSON(&validated); err != nil {
		if validationErrors := Validator.GetValidationErrors(err); validationErrors != nil {
			context.JSON(http.StatusBadRequest, gin.H{"message": validationError, "errors": validationErrors})
			return
		}
	}

	user, err := UserServices.CheckAuth(validated.Email, validated.Password)
	if err != nil {
		context.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": loginSuccessful, "data": user})
	return
}

func CreateAccount(context *gin.Context) {
	var validated Requests.CreateUserRequest
	if err := context.ShouldBindJSON(&validated); err != nil {
		if validationErrors := Validator.GetValidationErrors(err); validationErrors != nil {
			context.JSON(http.StatusBadRequest, gin.H{"message": validationError, "errors": validationErrors})
			return
		}
	}

	validated.Password, _ = UserServices.HashPassword(validated.Password)

	fmt.Println("hashed password is: ", validated.Password)

	user, err := Models.CreateNewUser(validated)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": createUserFailed, "data": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": user})
	return
}
