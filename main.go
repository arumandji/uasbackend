package main

import (
	"log"
	"os"

	"uas_backend/config"
	"uas_backend/database"
	"uas_backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load env
	config.LoadEnv()

	// Connect DB
	database.ConnectPostgres()
	database.ConnectMongo()

	// Router
	r := gin.Default()

	// Routes
	routes.RegisterRoutes(r)

	// Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running on port", port)
	r.Run(":" + port)
}
