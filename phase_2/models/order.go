package models

import (
    "github.com/AbhijeetSingh1801/go-ecom/phase_2/db"
    "log"
)

type Order struct {
    ID        int
    ProductID int
    Quantity  int
}

func CreateOrder(productID int, quantity int) {
    _, err := db.DB.Exec("INSERT INTO orders(product_id, quantity) VALUES(?, ?)", productID, quantity)
    if err != nil {
        log.Fatal(err)
    }
    log.Println("✅ Order created for product:", productID)
}
