package config

import (
	"Assignment-2/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// make connection postgres
var (
	DB_HOST = "localhost"
	DB_PORT = "5432"
	DB_USER = "agusilaban"
	DB_PASS = "agus080709"
	DB_NAME = "assignment-2"
	db      *gorm.DB
	err     error
)

func StartDB() *gorm.DB {
	config := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connection to database: ", err)
	}

	db.Debug().AutoMigrate(&models.Item{}, &models.Order{})

	return db
}
