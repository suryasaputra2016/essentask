package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func OpenPostgres() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("connecting postgres database: %w", err)
	}

	dsn := os.Getenv("DSN")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("connecting postgres database: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("connecting postgres database: %w", err)
	}

	log.Println("Postgres database connected.")
	return db, nil
}

func ClosePostgres(db *sql.DB) error {
	return db.Close()
}
