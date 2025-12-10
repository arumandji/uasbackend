package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectPostgres initializes PostgreSQL connection using GORM
func ConnectPostgres() {
	dsn := os.Getenv("POSTGRES_DSN")

	if dsn == "" {
		log.Fatal("❌ Missing POSTGRES_DSN in .env")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to PostgreSQL:", err)
	}

	DB = db
	fmt.Println("✔️ PostgreSQL connected")
}