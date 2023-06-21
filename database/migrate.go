package database

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

var DB *gorm.DB

func ConnectDB() {

	dsn := "root:@tcp(127.0.0.1:3306)/budget_app?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("cannot connect to database $v")
	}

	err = db.AutoMigrate(&User{}, &Wallet{}, &Transaction{})

	if err != nil {
		return
	}

	DB = db

}

func MigrateDb(context *gin.Context) {
	ConnectDB()
	err := DB.AutoMigrate(&User{}, &User{})

	if err != nil {
		fmt.Println("Database migration failed", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Database migration failed", "data": err})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Database migration successful", "data": nil})
	return
}
