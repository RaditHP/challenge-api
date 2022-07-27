package routes

import (
	"challenge-api/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	route := gin.Default()
	route.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	route.GET("/customers", controllers.ListCustomers)
	route.GET("/customers/:id", controllers.FindCustomer)
	route.POST("/customers", controllers.CreateCustomer)
	route.PATCH("/customers/:id", controllers.UpdateCustomer)
	route.DELETE("/customers/:id", controllers.DeleteCustomer)
	// route.GET("/customers/:limit", controllers.FindCustomerLimit)

	//route for Sell table
	route.GET("/sell", controllers.ListSell)
	route.GET("/sell/:id", controllers.GetSell)
	route.POST("/sell", controllers.CreateSell)
	route.PATCH("/sell/:id", controllers.UpdateSell)
	route.DELETE("/sell/:id", controllers.DeleteSell)
	return route
}
