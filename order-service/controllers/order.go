package controllers

import (
	"net/http"
	"pesto-ecommerce/order-service/database"
	"pesto-ecommerce/order-service/models"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save order
	result := database.DB.Create(&order)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, order)
}

func GetOrder(c *gin.Context) {
	id := c.Param("id")
	var order models.Order

	result := database.DB.First(&order, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, order)
}
