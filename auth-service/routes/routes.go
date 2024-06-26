package routes

import (
	"net/http"
	"pesto-ecommerce/auth-service/controllers"
	"pesto-ecommerce/auth-service/middleware"

	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.GET("/verify-token", controllers.VerifyToken)

	protected := r.Group("/protected")
	protected.Use(middleware.AuthMiddleware())
	protected.GET("/profile", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "You are authorized"})
	})
}
