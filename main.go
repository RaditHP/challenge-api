package main

import (
	"challenge-api/models"
	"challenge-api/routes"
)

// @title Gin Swagger challenge api
// @version 1.0
// @description This is my first attempt implementing Swagger
// @termsOfService http://swagger.io/terms
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:9000
// @BasePath /
// @schemes http

func main() {
	db := models.SetupDB()
	db.AutoMigrate(&models.Customer{})
	db.AutoMigrate(&models.Sell{}).AddForeignKey("customerid", "customers(id)", "CASCADE", "RESTRICT")

	route := routes.SetupRoutes(db)
	route.Run(":9000")
}
