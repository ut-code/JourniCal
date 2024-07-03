package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
	"github.com/ut-code/JourniCal/backend/app/env"
)

var db *gorm.DB

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
		fmt.Println("If this is not run from docker compose, this is probably not expected")
	}
	if env.DSN == "" {
		log.Fatalln("DSN environment variable not found")
	}
	var err error
	db, err = gorm.Open(postgres.Open(env.DSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error opening postgres: %v", err)
	}
}

func InitDB(t ...any) *gorm.DB {
	for _, v := range t {
		// migrate T to database
		if err := db.AutoMigrate(v); err != nil {
			log.Fatalf("Failed to migrate database: %v", err)
		}
	}
	return db
}
