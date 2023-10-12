package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	DB *gorm.DB
}

var database DB

func InitDatabaseConnection() {
	var ConnectionMasterDB string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&",
		"root",                 // database user
		"1234567890",           // database password
		"localhost",            // database host
		"3306",                 // database port
		"project-os-container", // database name
	)

	db, err := gorm.Open(mysql.Open(ConnectionMasterDB), &gorm.Config{})

	if err != nil {
		log.Fatalf(fmt.Sprintf("[Database] failed to connect database: %s\n", err.Error()))
	}

	database.DB = db
}
