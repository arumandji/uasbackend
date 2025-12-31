package routes

import (
	"github.com/gin-gonic/gin"
	"uas_backend/handlers"
	"uas_backend/middleware"
)

type RouteConfig struct {
	MahasiswaHandler *handlers.MahasiswaHandler
	DosenHandler     *handlers.DosenHandler
}

func RegisterRoutes(r *gin.Engine, h RouteConfig) {

	v1 := r.Group("/api/v1")

	// ================= AUTH =================
	auth := v1.Group("/auth")
	{
		auth.POST("/login", handlers.LoginHandler)
		auth.POST("/logout", middleware.JWTAuth(), handlers.LogoutHandler)
		auth.POST("/refresh", middleware.JWTAuth(), handlers.RefreshHandler)
		auth.GET("/profile", middleware.JWTAuth(), handlers.ProfileHandler)
	}

	// ================= PROTECTED =================
	api := v1.Group("")
	api.Use(middleware.JWTAuth())
	{
		// ---------- USERS (ADMIN) ----------
		users := api.Group("/users")
		{
			users.GET("/", handlers.GetUsers)
			users.PUT("/:id", handlers.UpdateUser)
			users.DELETE("/:id", handlers.DeleteUser)
		}

		// ---------- ACHIEVEMENTS ----------
		achievements := api.Group("/achievements")
		{
			achievements.GET("/", handlers.GetAllAchievements)
			achievements.POST("/", handlers.CreateAchievement)
			achievements.GET("/me", handlers.GetMyAchievements)
			achievements.PUT("/:id", handlers.UpdateAchievement)
			achievements.DELETE("/:id", handlers.DeleteAchievement)
		}

		// ---------- MAHASISWA ----------
		mahasiswa := api.Group("/mahasiswa")
		{
			mahasiswa.GET("/:id", h.MahasiswaHandler.GetMahasiswaByID)
			mahasiswa.GET("/user/:user_id", h.MahasiswaHandler.GetMahasiswaByUserID)
			mahasiswa.GET("/dosen/:dosen_wali_id", h.MahasiswaHandler.GetMahasiswaByAdvisor)
			mahasiswa.POST("/", h.MahasiswaHandler.CreateMahasiswa)
			mahasiswa.PUT("/:id", h.MahasiswaHandler.UpdateMahasiswa)
			mahasiswa.DELETE("/:id", h.MahasiswaHandler.DeleteMahasiswa)

		}

		// ---------- DOSEN ----------
		dosen := api.Group("/dosen")
		{
			dosen.GET("/", h.DosenHandler.GetAllDosen)
			dosen.GET("/:id", h.DosenHandler.GetDosenByID)
			dosen.GET("/user/:user_id", h.DosenHandler.GetDosenByUserID)
			dosen.POST("/", h.DosenHandler.CreateDosen)
			dosen.PUT("/:id", h.DosenHandler.UpdateDosen)
			dosen.DELETE("/:id", h.DosenHandler.DeleteDosen)

		}
	}
}
