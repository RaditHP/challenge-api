package models

import (
	"time"
)

type Customer struct {
	ID            uint      `json:"id" gorm:"primary_key"`
	Name          string    `json:"customername"`
	ContactNumber string    `json:"customeraddress"`
	Address       string    `json:"customercontno"`
	TotalBuy      string    `json:"totalbuy"`
	CreatorID     string    `json:"creatorid"`
	Date          time.Time `json:"date"`
}
