package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// VerifyAchievement godoc
// @Summary Verify achievement
// @Description Verify an achievement submission
// @Tags Achievement
// @Security BearerAuth
// @Produce json
// @Param id path string true "Achievement ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/achievements/{id}/verify [post]
// VerifyAchievement handles POST /achievements/:id/verify
func VerifyAchievement(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Achievement verified (placeholder)",
	})
}

// RejectAchievement godoc
// @Summary Reject achievement
// @Description Reject an achievement submission
// @Tags Achievement
// @Security BearerAuth
// @Produce json
// @Param id path string true "Achievement ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/achievements/{id}/reject [post]
// RejectAchievement handles POST /achievements/:id/reject
func RejectAchievement(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Achievement rejected (placeholder)",
	})
}
