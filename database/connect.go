package database

import (
	"fmt"
	"go-auth/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("nerov:{Ricchio666!}@/go_auth"), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database!")
	}else {
		fmt.Println("Connected to DB ")
	}

	DB = connection
	connection.AutoMigrate(&models.User{})
}