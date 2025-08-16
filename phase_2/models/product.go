package models

import (
	"log"

	"github.com/AbhijeetSingh1801/go-ecom/phase_2/db"
)

type Product struct {
	ID    int
	Name  string
	Price float64
	Stock int
}

func GetProducts() []Product {
	rows, err := db.DB.Query("SELECT id, name, price, stock FROM products")

	if err != nil {
		log.Fatal("Failed to get products: ", err)
	}
	defer rows.Close()
	var products []Product

	for rows.Next() {
		var p Product
		rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock)
		products = append(products, p)
	}
	return products
}

func AddProduct(name string, price float64) {
	_, err := db.DB.Exec("INSERT INTO products (name, price, stock) VALUES (?, ?, ?)", name, price, 100)
	if err != nil {
		log.Fatal("Failed to add product: ", err)
	}
	log.Println("✅ Product added:", name)
}
