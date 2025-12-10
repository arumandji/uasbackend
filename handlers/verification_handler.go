package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// VerifyAchievement handles POST /achievements/:id/verify
func VerifyAchievement(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Achievement verified (placeholder)",
	})
}

// RejectAchievement handles POST /achievements/:id/reject
func RejectAchievement(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Achievement rejected (placeholder)",
	})
}
