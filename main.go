package main

import (
    "log"
    "os"

    "UAS/config"
    "UAS/database"
    authModule "UAS/modules/auth"
    achModule "UAS/modules/achievements"
    "UAS/middleware"

    "github.com/gin-gonic/gin"
)

func main() {
    config.Load() // load env

    database.ConnectPostgres()
    database.ConnectMongo()

    r := gin.Default()

    api := r.Group("/api/v1")
    {
        auth := api.Group("/auth")
        authModule.RegisterAuthRoutes(auth)

        ach := api.Group("/achievements")
        ach.Use(middleware.JWTAuth())
        achModule.RegisterAchievementRoutes(ach)
    }

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    log.Printf("server running on :%s", port)
    r.Run(":" + port)
}
