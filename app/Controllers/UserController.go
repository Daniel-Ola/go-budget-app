package Controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nielscript.com/budgetapp/api/app/Models"
	UserServices "nielscript.com/budgetapp/api/app/Services"
)

func User(context *gin.Context) {

	//var user Models.User

	claims, _ := UserServices.GetTokenClaims(context)

	user, err := Models.FindUserByField("user_name", claims.Username)

	if err != nil {
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User info returned", "data": user})

}
