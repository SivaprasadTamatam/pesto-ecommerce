package database

import (
	"pesto-ecommerce/auth-service/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() {
	var err error

	models.DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	models.DB.AutoMigrate(&models.User{})
}
