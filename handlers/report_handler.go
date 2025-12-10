package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetStatistics handles /reports/statistics
func GetStatistics(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Statistics data (placeholder)",
	})
}