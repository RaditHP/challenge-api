package controllers

import (
	"models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateSellInput struct {
	sellID           uint      `json:sellid`
	ProductID        string    `json:productid`
	PursePrice       float64   `json:purseprice`
	SellPrice        float64   `json:sellprice`
	Quantity         int       `json:sellquantity`
	TotalPrice       float64   `json:selltotalprice`
	WarrantyVoidDate string    `json:sellwarrantyvoiddate`
	SellerID         int       `json:sellerid`
	SellDate         time.Time `json:selldate`
	CustomerID       uint      `gorm:"column:customerid"`
}

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

//Create a sell data
func CreateSell(c *gin.Context) {
	var input CreateSellInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sell := models.Sell{ProductID: input.ProductID, PursePrice: input.PursePrice,
		SellPrice: input.SellPrice, Quantity: input.Quantity, TotalPrice: input.TotalPrice, WarrantyVoidDate: input.WarrantyVoidDate, SellerID: input.SellerID, CustomerID: input.CustomerID}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&sell)
	c.JSON(http.StatusOK, gin.H{"data": sell})
}

//Delete a Sell Data
func DeleteSell(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var sell models.Sell
	if err := db.Where("id = $1", c.Param("id")).First(&sell).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data Not Found!"})
		return
	}
	db.Delete(&sell)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
