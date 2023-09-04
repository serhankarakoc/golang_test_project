package database

import (
	"zatrano/internal/app/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func SetupDatabase() *gorm.DB {
	if db != nil {
		return db
	}

	cfg := LoadConfig()
	dsn := GetDBConnectionString(cfg)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	err = db.AutoMigrate(&models.Author{})
	if err != nil {
		panic("Failed to migrate database")
	}

	return db
}
