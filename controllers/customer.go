package controllers

import (
	"challenge-api/models"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateCustomerInput struct {
	Name          string `json:"customername"`
	ContactNumber string `json:"customercontno"`
	Address       string `json:"customeraddress"`
	TotalBuy      string `json:"totalbuy"`
	CreatorID     string `json:"creatorid"`
	Date          string `json:"date"`
}

type UpdateCustomerInput struct {
	Name          string `json:"customername"`
	ContactNumber string `json:"customercontno"`
	Address       string `json:"customeraddress"`
	TotalBuy      string `json:"totalbuy"`
	CreatorID     string `json:"creatorid"`
	Date          string `json:"date"`
}

//GET /customer gets all customers in the table
// @Summary show the list of all customers
// @Description get a list of all the customers
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /customers [get]
func ListCustomers(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var customers []models.Customer
	db.Find(&customers)
	// This is used to return selected columns
	// db.Select([]string{"Name", "Address", "ID"}).Find(&customers)

	c.JSON(http.StatusOK, gin.H{"data": customers})

}

// @Summary show the list of all customers v2
// @Description get a list of all the customers v2
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /customers [get]
func ListCustomersV2(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var customers []models.Customer
	db.Find(&customers)

	c.JSON(http.StatusOK, gin.H{"data": customers})
}

// Get one customer with particular ID
// @Description get a list of all the customers v2
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Param id path string true "search customer by id"
// @Router /customers/{id} [get]
func FindCustomer(c *gin.Context) {
	var customer models.Customer

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = $1", c.Param("id")).First(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": customer})
}

// Get customer table limit data returned

// func FindCustomerLimit(c *gin.Context) {
// 	var customers []models.Customer

// 	db := c.MustGet("db").(*gorm.DB)

// 	db.Limit(c.Param("limit")).Find(&customers)

// 	c.JSON(http.StatusOK, gin.H{"data": customers})
// }

// create a new Customer entry
// Create e new customer
// @Description Create e new customer
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Param customer body CreateCustomerInput true "Customer data JSON"
// @Router /customers [post]
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

//Update Customer data
// Update all customer data
// @Description update a customer data
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Param customer body UpdateCustomerInput true "Customer data JSON"
// @Router /customers/{id} [put]
func UpdateCustomerPut(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var customer models.Customer
	if err := db.Where("id = $1", c.Param("id")).First(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data Not Found!"})
		return
	}

	var input UpdateCustomerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data Not Found!"})
		return
	}

	currentDate := time.Now().Format("19-10-2021")
	date, _ := time.Parse(currentDate, input.Date)
	var updatedInput models.Customer
	updatedInput.Name = input.Name
	updatedInput.ContactNumber = input.ContactNumber
	updatedInput.Address = input.Address
	updatedInput.TotalBuy = input.TotalBuy
	updatedInput.CreatorID = input.CreatorID
	updatedInput.Date = date

	db.Model(&customer).Updates(updatedInput)
	c.JSON(http.StatusOK, gin.H{"data": customer})

}

//Update Customer data
// @Description update a customer data
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Param customer body UpdateCustomerInput true "Customer data JSON"
// @Router /customers/{id} [patch]
func UpdateCustomer(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var customer models.Customer
	if err := db.Where("id = $1", c.Param("id")).First(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data Not Found!"})
		return
	}

	var input UpdateCustomerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data Not Found!"})
		return
	}

	currentDate := time.Now().Format("19-10-2021")
	date, _ := time.Parse(currentDate, input.Date)
	var updatedInput models.Customer
	updatedInput.Name = input.Name
	updatedInput.ContactNumber = input.ContactNumber
	updatedInput.Address = input.Address
	updatedInput.TotalBuy = input.TotalBuy
	updatedInput.CreatorID = input.CreatorID
	updatedInput.Date = date

	db.Model(&customer).Updates(updatedInput)
	c.JSON(http.StatusOK, gin.H{"data": customer})

}

//Delete Customer data
// @Description delete a customer data by id
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Param id path string true "ID to delete"
// @Router /customers/{id} [delete]
func DeleteCustomer(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var customer models.Customer
	if err := db.Where("id = $1", c.Param("id")).First(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data Not Found!"})
		return
	}
	db.Delete(&customer)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
