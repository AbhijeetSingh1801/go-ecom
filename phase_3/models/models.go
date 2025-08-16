package models

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type CartItem struct {
	ID        int `json:"id"`
	ProductID int `json:"product_id"`
	Qty       int `json:"qty"`
}

type Order struct {
	ID        int    `json:"id"`
	ProductID int    `json:"product_id"`
	Qty       int    `json:"qty"`
	CreatedAt string `json:"created_at"`
}
