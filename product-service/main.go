package main

import (
	"log"
	"pesto-ecommerce/product-service/database"
	"pesto-ecommerce/product-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	// initialize the database
	database.Connect()
	database.Migrate()

	r := gin.Default()

	routes.Setup(r)

	if err := r.Run(":8081"); err != nil {
		log.Fatal("Failed to run server: %v", err)
	}
}
