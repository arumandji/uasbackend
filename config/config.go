package config

import (
    "github.com/joho/godotenv"
    "log"
)

func Load() {
    if err := godotenv.Load(); err != nil {
        log.Println(".env not found, using env variables")
    }
}
