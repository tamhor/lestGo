package database

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       int
	Uuid     string
	Username string
	Password string
	Email    string
}

type Article struct {
	gorm.Model
	Id      int
	Uuid    string
	Title   string
	Content string
}

func Migration() {
	ConnectDB()
	DB.AutoMigrate(&User{}, &Article{})
	fmt.Println("Migration Success")
}
