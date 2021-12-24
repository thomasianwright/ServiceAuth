package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"serviceauth/models"
)

var DB *gorm.DB

func Setup() {
	db, err := gorm.Open(mysql.Open(os.Getenv("mysql")), &gorm.Config{})

	DB = db

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Application{}, &models.Scope{})
}
