package main

import (
	"challenge-api/models"
	"challenge-api/routes"
)

func main() {
	db := models.SetupDB()
	db.AutoMigrate(&models.Customer{})
	db.AutoMigrate(&models.Sell{}).AddForeignKey("customerid", "customers(id)", "CASCADE", "RESTRICT")

	route := routes.SetupRoutes(db)
	route.Run()
}
