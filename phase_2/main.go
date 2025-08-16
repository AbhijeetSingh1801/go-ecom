package main

import (
	"log"
	"github.com/AbhijeetSingh1801/go-ecom/phase_2/db"
	"github.com/AbhijeetSingh1801/go-ecom/phase_2/models"
)

func main() {
	db.Connect()
	db.Migrate()

	// seed product
	models.AddProduct("Laptop", 999.99)
	models.AddProduct("Phone", 499.50)
  
	// show products
	products := models.GetProducts()
	log.Println("📦 Products:", products)

	models.CreateOrder(1, 2)
}