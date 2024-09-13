package models

import (
	"github.com/jinzhu/gorm"
)

// Product represents a product in the inventory
type Product struct {
	gorm.Model
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

// Migrate the schema (create table if not exists)
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Product{})
}
