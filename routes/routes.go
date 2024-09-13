package routes

import (
	controllers "go-product-api/controller"
	"go-product-api/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// SetupRouter initializes the routes and DB middleware
func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// Middleware to attach DB connection to the context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// Product routes
	r.GET("/products", controllers.GetProducts)
	r.GET("/products/:id", controllers.GetProduct)
	r.POST("/products", controllers.CreateProduct)
	r.PUT("/products/:id", controllers.UpdateProduct)
	r.DELETE("/products/:id", controllers.DeleteProduct)

	// Run DB migrations
	models.Migrate(db)

	return r
}
