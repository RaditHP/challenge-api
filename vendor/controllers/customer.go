package controllers

import (
	"models"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func ListCustomers(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var customers []models.Customer
	db.Find(&customers)

	c.JSON(http.StatusOK, gin.H{"data": customers})
}
