package main

import (
	"challenge-api/controllers"
	"challenge-api/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	return router
}

func TestGetCustomer(t *testing.T) {
	db := models.SetupDB()
	r := SetupRouter(db)
	r.GET("/customers", controllers.ListCustomers)
	req, _ := http.NewRequest("GET", "/customers", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var customer models.CustResp
	json.Unmarshal(w.Body.Bytes(), &customer)

	assert.Equal(t, http.StatusOK, w.Code, "Get list of customer")

	assert.Equal(t, 6, len(customer.Data), "Get length")
}

func TestFindCustomer(t *testing.T) {
	db := models.SetupDB()
	r := SetupRouter(db)
	r.GET("/customers/:id", controllers.FindCustomer)
	req, _ := http.NewRequest("GET", "/customers/2", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var customer models.CustSingleResp

	mockCustomer := models.CustSingleResp{
		Data: models.Customer{ID: 2,
			Name:          "JB",
			ContactNumber: "5142",
			Address:       "Depok",
			TotalBuy:      "4",
			CreatorID:     "5",
			Date:          time.Date(2022, time.September, 9, 17, 0, 0, 0, time.UTC),
			Sells:         []models.Sell(nil)},
	}

	json.Unmarshal(w.Body.Bytes(), &customer)

	assert.Equal(t, http.StatusOK, w.Code, "Get single customer")

	assert.Equal(t, mockCustomer, customer, "check struct contents")
}

func TestGetCustomerV11(t *testing.T) {
	db := models.SetupDB()
	r := SetupRouter(db)
	r.GET("/v1.1/customers", controllers.ListCustomersV2)
	req, _ := http.NewRequest("GET", "/v1.1/customers", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var customer models.CustResp
	json.Unmarshal(w.Body.Bytes(), &customer)

	assert.Equal(t, http.StatusOK, w.Code, "Get list of customer")

	assert.Equal(t, 6, len(customer.Data), "Get length")
}
