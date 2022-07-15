package main

import (
	"models"
	"routes"
)

func main() {
	db := models.SetupDB()
	db.AutoMigrate(&models.Customer{})
	db.AutoMigrate(&models.Sell{}).AddForeignKey("customerid", "customers(id)", "CASCADE", "RESTRICT")

	// db.Migrator().CreateConstraint(&models.Sell{}, "Customers")
	// db.Migrator().CreateConstraint(&models.Sell{}, "fk_customers_sells")

	route := routes.SetupRoutes(db)
	route.Run()
}
