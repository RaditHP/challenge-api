package routes

import (
	"challenge-api/controllers"

	_ "challenge-api/docs"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	route := gin.Default()
	url := ginSwagger.URL("http://localhost:9000/swagger/doc.json")
	route.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	route.GET("/customers", controllers.ListCustomers)
	route.GET("/customers/:id", controllers.FindCustomer)
	route.POST("/customers", controllers.CreateCustomer)
	route.PUT("/customers/:id", controllers.UpdateCustomerPut)
	route.PATCH("/customers/:id", controllers.UpdateCustomer)
	route.DELETE("/customers/:id", controllers.DeleteCustomer)
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	// route.GET("/customers/:limit", controllers.FindCustomerLimit)

	//route for Sell table
	route.GET("/sell", controllers.ListSell)
	route.GET("/sell/:id", controllers.GetSell)
	route.POST("/sell", controllers.CreateSell)
	route.PATCH("/sell/:id", controllers.UpdateSell)
	route.DELETE("/sell/:id", controllers.DeleteSell)

	v11 := route.Group("/v1.1")
	{
		v11.GET("/customers", controllers.ListCustomersV2)
	}
	return route
}
