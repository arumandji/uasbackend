package main

import (
	"log"
	"os"

	"uas_backend/config"
	"uas_backend/database"
	"uas_backend/routes"

	"uas_backend/handlers"
	"uas_backend/repositories"
	"uas_backend/services"

	_ "uas_backend/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title UAS Backend API
// @version 1.0
// @description API Sistem Prestasi Mahasiswa
// @contact.name Arum
// @contact.email arum@mail.com

// @host localhost:8080
// @BasePath /api
// @schemes http

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {

	config.LoadEnv()

	database.ConnectPostgres()
	db := database.DB
	database.ConnectMongo()

	// ================= REPOSITORY =================
	mahasiswaRepo := repositories.NewMahasiswaRepository(db)
	dosenRepo := repositories.NewDosenRepository(db)

	// ================= SERVICE =================
	mahasiswaService := services.NewMahasiswaService(mahasiswaRepo)
	dosenService := services.NewDosenService(dosenRepo)

	// ================= HANDLER =================
	mahasiswaHandler := handlers.NewMahasiswaHandler(mahasiswaService)
	dosenHandler := handlers.NewDosenHandler(dosenService)

	r := gin.Default()

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// ================= ROUTES =================
	routes.RegisterRoutes(r, routes.RouteConfig{
		MahasiswaHandler: mahasiswaHandler,
		DosenHandler:     dosenHandler,
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("\n\033[32m=================================================\033[0m")
	log.Println("\033[32m✅  SEMUA DATABASE TERHUBUNG!\033[0m")
	log.Println("\033[32m🚀  SERVER SIAP DI PORT :" + port + "\033[0m")
	log.Println("\033[32m👉  Swagger: http://localhost:" + port + "/swagger/index.html\033[0m")
	log.Println("\033[32m=================================================\033[0m\n")

	r.Run(":" + port)
}
