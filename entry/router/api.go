package router

import (
	"github.com/gin-gonic/gin"
	"nielscript.com/budgetapp/api/app/Controllers"
	"nielscript.com/budgetapp/api/database"
	"nielscript.com/budgetapp/api/entry/middlewares"
)

func Routes() {

	r := gin.Default()

	authRoutes := r.Group("/auth")

	{
		authRoutes.POST("/register", Controllers.CreateAccount)
		authRoutes.POST("/login", Controllers.Login)
	}

	userRoutes := r.Group("/user").Use(middlewares.AuthRequired())

	{
		userRoutes.GET("/", Controllers.User)
	}

	utilities := r.Group("/utilities")

	{
		utilities.GET("/migrate-db", database.MigrateDb)
	}

	r.Run("localhost:8010") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
