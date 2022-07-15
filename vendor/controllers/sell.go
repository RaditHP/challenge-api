package controllers

import (
	"models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// List all the data in the sell table
func ListSell(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var sells []models.Sell
	db.Find(&sells)

	c.JSON(http.StatusOK, gin.H{"data": sells})
}

// Get one sell where the id for the sell is requested
func GetSell(c *gin.Context) {
	var sell models.Sell

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = $1", c.Param("id")).First(&sell).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": sell})
}
