package database

import (
    "log"
    "os"

    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq"
)

var DB *sqlx.DB

func ConnectPostgres() {
    dsn := os.Getenv("POSTGRES_DSN")
    db, err := sqlx.Connect("postgres", dsn)
    if err != nil {
        log.Fatalf("postgres connect err: %v", err)
    }
    DB = db
}