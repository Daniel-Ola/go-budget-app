package main

import (
	"nielscript.com/budgetapp/api/database"
	"nielscript.com/budgetapp/api/entry/router"
)

func main() {
	//r := gin.Default()

	database.ConnectDB()

	router.Routes()

}
