package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

var dsn string
var db *gorm.DB

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
		fmt.Println("If this is not run from docker compose, this is probably not expected")
	}
	dsn = os.Getenv("DSN")
	if dsn == "" {
		log.Fatalln("DSN environment variable not found")
	}
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error opening postgres: %v", err)
	}
}

func InitDB(t any) *gorm.DB {
	// migrate T to database
	if err := db.AutoMigrate(t); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	return db
}
