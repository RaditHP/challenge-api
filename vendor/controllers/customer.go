package controllers

import (
	"models"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateCustomerInput struct {
	Name          string `json:"customername"`
	ContactNumber string `json:"customeraddress"`
	Address       string `json:"customercontno"`
	TotalBuy      string `json:"totalbuy"`
	CreatorID     string `json:"creatorid"`
	Date          string `json:"date"`
}

//GET /customer gets all customers in the table
func ListCustomers(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var customers []models.Customer
	db.Find(&customers)

	c.JSON(http.StatusOK, gin.H{"data": customers})
}

// Get one customer with particular ID
func FindCustomer(c *gin.Context) {
	var customer models.Customer

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = $1", c.Param("id")).First(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": customer})
}

func CreateCustomer(c *gin.Context) {
	var input CreateCustomerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	currentDate := time.Now().Format("19-10-2021")
	date, _ := time.Parse(currentDate, input.Date)

	//create customer
	customer := models.Customer{Name: input.Name, ContactNumber: input.ContactNumber,
		Address: input.Address, TotalBuy: input.TotalBuy, CreatorID: input.CreatorID, Date: date}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&customer)

	c.JSON(http.StatusOK, gin.H{"data": customer})
}
