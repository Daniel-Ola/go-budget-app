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

	token, err := UserServices.GenerateJWT(user.UserName, context.GetHeader("app-key"))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Error " + err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": loginSuccessful, "data": user, "token": token})
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

// payload
/*

username
email
names
refresh token (encrypted)
expire
stillValid
signature - app-keys

when user signs in and credentials verified
a row is created - userId, timestamps, stillValid, refreshToken, lastUsed
the refresh token is encrypted and sent together with the JWT which is just created
when user signs out, the token is deleted and the jWT expires

*/
