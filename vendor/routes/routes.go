package routes

import (
	"controllers"

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
	return route
}
