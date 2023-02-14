package models

import (
	"time"
)

type Customer struct {
	ID            uint      `json:"id" gorm:"primary_key"`
	Name          string    `json:"customername"`
	ContactNumber string    `json:"customercontno"`
	Address       string    `json:"customeraddress"`
	TotalBuy      string    `json:"totalbuy"`
	CreatorID     string    `json:"creatorid"`
	Date          time.Time `json:"date"`
	Sells         []Sell
}

type CustResp struct {
	Data []Customer `json:"data"`
}

type CustSingleResp struct {
	Data Customer `json:"data"`
}
