package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	UserServices "nielscript.com/budgetapp/api/app/Services"
	"time"
)

var authRequired = "Authentication is required"

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {

		authToken := c.GetHeader("auth-token")
		if len(authToken) < 1 {
			c.JSON(http.StatusUnauthorized, gin.H{"message": authRequired})
			c.Abort()
			return
		}

		_, err := UserServices.VerifyToken(authToken, c.GetHeader("app-key"))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			c.Abort()
			return
		}

		/*if isValid {
			c.Next()
		}*/

		c.Next()

		//c.AbortWithError(http.StatusUnauthorized, gin.H{"message": "You're not signed in"})

		// after request
	}
}

func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}
