package controllers

import (
	"go-product-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Get all products
func GetProducts(c *gin.Context) {
	var products []models.Product
	db := c.MustGet("db").(*gorm.DB)
	db.Find(&products)
	c.JSON(http.StatusOK, products)
}

// Get a single product by ID
func GetProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	db := c.MustGet("db").(*gorm.DB)
	if db.First(&product, id).RecordNotFound() {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

// Create a new product
func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&product)
	c.JSON(http.StatusCreated, product)
}

// Update an existing product
func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	db := c.MustGet("db").(*gorm.DB)
	if db.First(&product, id).RecordNotFound() {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Save(&product)
	c.JSON(http.StatusOK, product)
}

// Delete a product
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	db := c.MustGet("db").(*gorm.DB)
	if db.First(&product, id).RecordNotFound() {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	db.Delete(&product)
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}
