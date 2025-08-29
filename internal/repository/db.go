package repository

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const connStr = "user=postgres dbname=taas password=root sslmode=disable"

func InitDB() (*sql.DB, error) {

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err := createTable(db); err != nil {
		panic(err)
	}

	return db, nil
}

func createTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username TEXT NOT NULL UNIQUE,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`

	_, err := db.Exec(query)

	if err != nil {
		return fmt.Errorf("failed to create users table: %w", err)
	}
	fmt.Println("Users table created/migrated successfully")
	return nil
}
