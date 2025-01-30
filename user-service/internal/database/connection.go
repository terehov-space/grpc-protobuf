package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"wpc/user-service/internal/models"
)

func InitDB() *gorm.DB {
	dsn := "postgresql://ovr_user:ovr_pass@localhost:5429/performance"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Failed to automigrate: %v\n", err)
	}

	return db
}
