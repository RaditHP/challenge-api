package main

import (
	"models"
	"routes"
)

func main() {
	db := models.SetupDB()
	db.AutoMigrate(&models.Customer{})

	route := routes.SetupRoutes(db)
	route.Run()
}
