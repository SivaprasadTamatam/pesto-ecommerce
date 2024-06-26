package main

import (
	"log"
	"pesto-ecommerce/order-service/database"
	"pesto-ecommerce/order-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	database.Connect()
	database.Migrate()

	r := gin.Default()
	routes.Setup(r)

	if err := r.Run(":8082"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
