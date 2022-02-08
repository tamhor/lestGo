package database

import (
	"fmt"
	"log"
	"strconv"

	"github.com/tamhor/lestGo/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	db_conn := config.Config("DB_CONNECTION")

	if db_conn == "postgres" {
		p := config.Config("DB_PORT")
		port, err := strconv.ParseUint(p, 10, 32)

		if err != nil {
			log.Println("Please check yout config")
		}
		conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))
		DB, err = gorm.Open(postgres.Open(conn))
		if err != nil {
			panic("Failed to connect database")
		}
		fmt.Println("Connection Opened to Postgree Database")
	} else if db_conn == "sqlite" {
		DB, err = gorm.Open(sqlite.Open("database/gorm.db"))
		if err != nil {
			panic("Failed to connect database")
		}
		fmt.Println("Connection Opened to SQLite Database")
	} else if db_conn == "mysql" {
		p := config.Config("DB_PORT")
		port, err := strconv.ParseUint(p, 10, 32)

		if err != nil {
			log.Println("Please check yout config")
		}
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_HOST"), port, config.Config("DB_NAME"))
		DB, err = gorm.Open(mysql.Open(dsn))
		if err != nil {
			panic("Failed to connect database")
		}
		fmt.Println("Connection Opened to MySQL Database")
	} else {
		panic("No Setting")
	}

}
