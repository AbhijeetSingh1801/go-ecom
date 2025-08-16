package db

import (
	"database/sql"
	"log"
	"os"
	_ "github.com/mattn/go-sqlite3"
)


func InitDB() *sql.DB {
	dbPath := os.Getenv("DB_PATH")

	if dbPath == "" {
		log.Fatal("DB_PATH is not set")
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal("Failed to connect to DB: ", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("DB unreachable: ", err)
	}

	// Create tables
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS products (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			price REAL
		);
		CREATE TABLE IF NOT EXISTS cart (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			product_id INTEGER,
			qty INTEGER
		);
		CREATE TABLE IF NOT EXISTS orders (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			product_id INTEGER,
			qty INTEGER,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`)
	if err != nil {
		log.Fatal("❌ Failed to create tables:", err)
	}

	return db
}