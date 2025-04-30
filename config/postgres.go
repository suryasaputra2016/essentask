package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var dsn = fmt.Sprintf(
	"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
	os.Getenv("POSTGRES_HOST"),
	os.Getenv("POSTGRES_PORT"),
	os.Getenv("POSTGRES_USER"),
	os.Getenv("POSTGRES_PASSWORD"),
	os.Getenv("POSTGRES_DB"),
	"disable",
)

func OpenPostgres() (*sql.DB, error) {

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
