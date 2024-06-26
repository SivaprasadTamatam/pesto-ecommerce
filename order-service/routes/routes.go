package routes

import (
	"pesto-ecommerce/order-service/controllers"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {
	router.POST("/orders", controllers.CreateOrder)
	router.GET("/orders/:id", controllers.GetOrder)
}
