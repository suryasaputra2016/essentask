package config

import (
	"database/sql"
	"fmt"
	"log"
)

func PrepareTables(db *sql.DB) error {
	err := addUserTable(db)
	if err != nil {
		return fmt.Errorf("preparing tables: %w", err)
	}
	err = addTaskTable(db)
	if err != nil {
		return fmt.Errorf("preparing tables: %w", err)
	}
	log.Println("all tables created or already existed")
	return nil
}

func addUserTable(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name VARCHAR(50),
			email VARCHAR(100) UNIQUE NOT NULL,
			password_hash TEXT NOT NULL,
			is_verified BOOL DEFAULT false,
			role VARCHAR(5) DEFAULT 'user'
		);`
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("adding user table: %w", err)
	}
	log.Println("user table created or already existed")
	return nil
}

func addTaskTable(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS tasks(
			id SERIAL PRIMARY KEY,
			user_id INT REFERENCES users(id) NOT NULL,
			name VARCHAR(50) NOT NULL,
			status VARCHAR(10) DEFAULT 'planned',
			deadline TIME
		);`
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("adding task table: %w", err)
	}
	log.Println("task table created or already existed")
	return nil
}
