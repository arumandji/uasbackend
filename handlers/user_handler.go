package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUsers returns all users (starter)
func GetUsers(c *gin.Context) {
	// TODO: integrate with user_service
	c.JSON(http.StatusOK, gin.H{
		"message": "Get all users (placeholder)",
	})
}