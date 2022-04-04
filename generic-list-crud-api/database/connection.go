package database

import (
	"fmt"

	"generic-list-crud-api/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	conn, err := gorm.Open(mysql.Open("root:root@/db?parseTime=true"), &gorm.Config{})

	if err != nil {
		panic(fmt.Sprintf("Could not connect to the database, Error: %v", err))
	}

	DB = conn

	err = conn.AutoMigrate(&models.List{}, &models.Symbol{})

	if err != nil {
		panic(fmt.Sprintf("Could not migrate to the database, Error: %v", err))
	}
}
