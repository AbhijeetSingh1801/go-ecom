package db

import "log"

func Migrate() {

	productTable := `
    CREATE TABLE IF NOT EXISTS products (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        price REAL NOT NULL,
        stock INTEGER NOT NULL
    );`

	orderTable := `
    CREATE TABLE IF NOT EXISTS orders (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        product_id INTEGER,
        quantity INTEGER,
        FOREIGN KEY(product_id) REFERENCES products(id)
    );`

	_, err := DB.Exec(productTable)
	if err != nil {
		log.Fatal("Failed to create product table: ", err)
	}

	_, err = DB.Exec(orderTable)
	if err != nil {
		log.Fatal("Failed to create order table: ", err)
	}

	log.Println("Database migrated successfully")
}
