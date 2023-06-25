package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"nielscript.com/budgetapp/api/app/Models"
	"nielscript.com/budgetapp/api/app/Requests"
	UserServices "nielscript.com/budgetapp/api/app/Services"
	"nielscript.com/budgetapp/api/app/Validator"
)

const loginSuccessful, validationError = "User signed in successfully", "Validation Error"
const createUserFailed = "Failed to create user"

func Login(context *gin.Context) {
	var validated Requests.CreateUserRequest
	if err := context.ShouldBindJSON(&validated); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"firstName": validated.FirstName})

	//params := context.Request.Body
	//fmt.Println(params)
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
