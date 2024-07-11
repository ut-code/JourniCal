package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/ut-code/JourniCal/backend/app/env/secret"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(postgres.Open(secret.DSN), &gorm.Config{})
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
