package config

import (
	"fmt"
	"log"

	"github.com/Sunwatcha303/Project-OS-Container/constants"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	DB *gorm.DB
}

var Database DB

func InitDatabaseConnection() {
	var ConnectionMasterDB string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true",
		constants.Database_user,
		constants.Database_password,
		constants.Database_host_for_test,
		constants.Database_port,
		constants.Database_name,
	)

	db, err := gorm.Open(mysql.Open(ConnectionMasterDB), &gorm.Config{})

	if err != nil {
		log.Fatalf(fmt.Sprintf("[Database] failed to connect database: %s\n", err.Error()))
	}

	Database.DB = db
}
