package main

import (
	"pesto-ecommerce/auth-service/database"
	"pesto-ecommerce/auth-service/routes"

	"github.com/gin-gonic/gin"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var err error

func main() {

	// Initialize DB
	database.Connect()

	// Initialize the Gin router
	r := gin.Default()

	// Setup routers
	routes.Setup(r)

	r.Run((":8080"))
}
