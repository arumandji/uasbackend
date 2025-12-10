package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateAchievement handles POST /achievements
func CreateAchievement(c *gin.Context) {
	// TODO: bind JSON & call service

	c.JSON(http.StatusOK, gin.H{
		"message": "Achievement created (placeholder)",
	})
}

// GetAllAchievements handles GET /achievements
func GetAllAchievements(c *gin.Context) {
	// TODO: fetch achievements from service

	c.JSON(http.StatusOK, gin.H{
		"message": "List achievements (placeholder)",
	})
}
