package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var db *gorm.DB

// ConnectDatabase connects to the PostgreSQL database
func ConnectDatabase() (*gorm.DB, error) {
	host := os.Getenv("HOSTNAME")
	port := os.Getenv("PORT")
	user := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	dbname := "postgres"

	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)
	fmt.Println(psqlInfo)

	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		log.Println("Failed to connect to database", err)
		return nil, err
	}

	return db, nil
}
