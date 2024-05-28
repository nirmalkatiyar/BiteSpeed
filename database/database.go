package database

import (
	"fmt"
	"github.com/nirmalkatiyar/bitespeed/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

// InitDB ...make connection with db
func InitDB() *gorm.DB {
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&models.Contact{})
	if err != nil {
		panic("failed to migrate database")
		return nil
	}
	return db
}
