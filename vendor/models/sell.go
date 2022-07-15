package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Sell struct {
	gorm.Model
	//ID               uint      `json:id gorm:"primary_key"`
	sellID           int       `json:sellid`
	ProductID        string    `json:productid`
	PursePrice       float64   `json:purseprice`
	SellPrice        float64   `json:sellprice`
	Quantity         int       `json:sellquantity`
	TotalPrice       float64   `json:selltotalprice`
	WarrantyVoidDate string    `json:sellwarrantyvoiddate`
	SellerID         int       `json:sellerid`
	SellDate         time.Time `json:selldate`
	CustomerID       uint      `gorm:"column:customerid"`
	//Customer         Customer
}
