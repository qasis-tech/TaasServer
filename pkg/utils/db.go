package utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, falling back to system env")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// First connect to default "postgres"
	defaultConnStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=postgres sslmode=disable",
		host, port, user, password)

	defaultDB, err := sql.Open("postgres", defaultConnStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to default DB: %w", err)
	}
	defer defaultDB.Close()

	var exists bool
	err = defaultDB.QueryRow("SELECT 1 FROM pg_database WHERE datname=$1", dbname).Scan(&exists)
	if err == sql.ErrNoRows {
		_, err = defaultDB.Exec("CREATE DATABASE " + dbname)
		if err != nil {
			return nil, fmt.Errorf("failed to create database: %w", err)
		}
		fmt.Println("✅ Database created successfully:", dbname)
	} else if err != nil {
		return nil, fmt.Errorf("failed to check database existence: %w", err)
	}

	// Connect to actual target DB
	targetConnStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", targetConnStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to target DB: %w", err)
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
		ProfilePic TEXT DEFAULT '',
		roles TEXT NOT NULL,
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
