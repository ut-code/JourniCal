package helper

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/ut-code/JourniCal/backend/types"

	"github.com/joho/godotenv"
)

var Database *gorm.DB

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dsn := os.Getenv("DSN")
	var err error
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error opening postgres: %v", err)
	}
	// migrate Diary to database
	if err := Database.AutoMigrate(&types.Diary{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}
