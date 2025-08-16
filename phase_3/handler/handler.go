package handler

import (
	"database/sql"
	"net/http"

	"github.com/AbhijeetSingh1801/go-ecom/phase_3/models"
	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context, db *sql.DB) {
	rows, _ := db.Query("SELECT id, name, price FROM products")
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		rows.Scan(&p.ID, &p.Name, &p.Price)
		products = append(products, p)
	}

	c.JSON(http.StatusOK, products)
}

func GetCart(c *gin.Context, db *sql.DB) {
	rows, _ := db.Query("SELECT id, product_id, qty FROM cart")
	defer rows.Close()

	var cart []models.CartItem
	for rows.Next() {
		var c models.CartItem
		rows.Scan(&c.ID, &c.ProductID, &c.Qty)
		cart = append(cart, c)
	}

	c.JSON(http.StatusOK, cart)
}

func AddToCart(c *gin.Context, db *sql.DB) {
	var item models.CartItem

	if err := c.ShouldBindJSON(&item); err != nil || item.Qty <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Exec("INSERT INTO cart (product_id, qty) VALUES (?, ?)", item.ProductID, item.Qty)
	c.JSON(http.StatusOK, gin.H{"message": "Item added to cart"})

}
func RemoveFromCart(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	db.Exec("DELETE FROM cart WHERE id = ?", id)
	c.JSON(http.StatusOK, gin.H{"status": "item removed"})
}

// Checkout
func Checkout(c *gin.Context, db *sql.DB) {
	// Move cart items to orders
	_, _ = db.Exec("INSERT INTO orders (product_id, qty) SELECT product_id, qty FROM cart")
	_, _ = db.Exec("DELETE FROM cart")
	c.JSON(http.StatusOK, gin.H{"status": "checkout complete"})
}

// Orders
func OrderHistory(c *gin.Context, db *sql.DB) {
	rows, _ := db.Query("SELECT id, product_id, qty, created_at FROM orders ORDER BY created_at DESC")
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var o models.Order
		rows.Scan(&o.ID, &o.ProductID, &o.Qty, &o.CreatedAt)
		orders = append(orders, o)
	}
	c.JSON(http.StatusOK, orders)
}
