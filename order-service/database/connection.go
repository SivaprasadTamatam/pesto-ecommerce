package database

import (
	"log"
	"pesto-ecommerce/order-service/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
}

func Migrate() {
	err := DB.AutoMigrate(&models.Order{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}
