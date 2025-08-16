package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/AbhijeetSingh1801/go-ecom/phase_3/db"
	"github.com/AbhijeetSingh1801/go-ecom/phase_3/models"
	"github.com/AbhijeetSingh1801/go-ecom/phase_3/routes"
	"github.com/AbhijeetSingh1801/go-ecom/phase_3/handler"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ No .env file found, using defaults")
	}

	db := db.InitDB()
	defer db.Close()
	models.SeedProducts(db)

	r := gin.Default()
	r.Use(handler.LoggerMiddleware())

	routes.SetupRoutes(r, db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
