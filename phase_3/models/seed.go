package models

import (
	"database/sql"
)


func SeedProducts(db *sql.DB) {

	var count int 
	db.QueryRow("SELECT COUNT(*) FROM products").Scan(&count)

	if count > 0 {
		return
	}

	// Insert demo products
	db.Exec("INSERT INTO products (name, price) VALUES (?, ?)", "Laptop", 75000)
	db.Exec("INSERT INTO products (name, price) VALUES (?, ?)", "Phone", 30000)
	db.Exec("INSERT INTO products (name, price) VALUES (?, ?)", "Headphones", 2000)
}