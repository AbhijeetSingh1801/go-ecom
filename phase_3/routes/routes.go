package routes

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/AbhijeetSingh1801/go-ecom/phase_3/handler"
)

func SetupRoutes(r *gin.Engine, db *sql.DB) {
	r.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"status": "ok"}) })

	r.GET("/products", func(c *gin.Context) { handler.GetProducts(c, db) })

	r.GET("/cart", func(c *gin.Context) { handler.GetCart(c, db) })
	r.POST("/cart/add", func(c *gin.Context) { handler.AddToCart(c, db) })
	r.DELETE("/cart/:id", func(c *gin.Context) { handler.RemoveFromCart(c, db) })

	r.POST("/checkout", func(c *gin.Context) { handler.Checkout(c, db) })
	r.GET("/orders/history", func(c *gin.Context) { handler.OrderHistory(c, db) })
}
