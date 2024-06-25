package routes

import (
	"pesto-ecommerce/product-service/controllers"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {
	router.GET("/products", controllers.GetProducts)
	router.GET("/products/:id", controllers.GetProduct)
	router.POST("/products", controllers.CreateProduct)
	router.PUT("/products/:id", controllers.UpdateProduct)
	router.DELETE("/products/:id", controllers.DeleteProduct)
}
