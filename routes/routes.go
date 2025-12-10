package routes

import (
	"github.com/gin-gonic/gin"
	"uas_backend/handlers"
	"uas_backend/middleware"
)

func RegisterRoutes(r *gin.Engine) {

	// Public
	auth := r.Group("/api/auth")
	{
		auth.POST("/login", handlers.LoginHandler)
	}

	// Protected
	api := r.Group("/api")
	api.Use(middleware.JWTAuth())
	{
		api.GET("/profile", handlers.ProfileHandler)

		user := api.Group("/users")
		{
			user.GET("/", handlers.GetUsers)
		}

		achievement := api.Group("/achievements")
		{
			achievement.GET("/", handlers.GetAllAchievements)
			achievement.POST("/", handlers.CreateAchievement)
		}
	}
}
