package config

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadEnv loads variables from .env file
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ .env file not found, using system environment variables...")
	}
}