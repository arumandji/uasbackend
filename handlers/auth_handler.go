package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginHandler handles login requests (starter version)
func LoginHandler(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// TODO: integrate with auth_service
	c.JSON(http.StatusOK, gin.H{
		"message": "Login success (placeholder)",
		"user":    req.Username,
		"token":   "dummy-token",
	})
}

// ProfileHandler returns dummy user profile
func ProfileHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Profile OK",
		"user":    "Placeholder User",
	})
}