package handlers

import (
	"fmt"
	"net/http"

	"uas_backend/database"
	"uas_backend/models"
	"uas_backend/middleware"

	"github.com/gin-gonic/gin"
)

// LoginHandler godoc
// @Summary Login user
// @Description Login dan mendapatkan JWT token
// @Tags Auth
// @Accept multipart/form-data
// @Produce json
// @Param username formData string true "Username"
// @Param password formData string true "Password"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /v1/auth/login [post]
func LoginHandler(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username & password required"})
		return
	}

	var user models.User
	err := database.DB.
		Where("username = ?", username).
		First(&user).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	// sementara TANPA hashing
	if user.Password != password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid password"})
		return
	}

	token, err := middleware.GenerateToken(
		fmt.Sprint(user.ID),
		user.RoleID,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "login success",
		"token":   token,
	})
}

// ProfileHandler godoc
// @Summary Get profile
// @Description Get current user profile
// @Tags Auth
// @Security BearerAuth
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]string
// @Router /v1/auth/profile [get]
func ProfileHandler(c *gin.Context) {
	userID, _ := c.Get("user_id")
	role, _ := c.Get("role")

	c.JSON(http.StatusOK, gin.H{
		"user_id": userID,
		"role":    role,
	})
}

// LogoutHandler godoc
// @Summary Logout
// @Tags Auth
// @Security BearerAuth
// @Produce json
// @Success 200 {object} map[string]string
// @Router /v1/auth/logout [post]
func LogoutHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "logout success"})
}

// RefreshHandler godoc
// @Summary Refresh token
// @Tags Auth
// @Security BearerAuth
// @Produce json
// @Success 200 {object} map[string]string
// @Router /v1/auth/refresh [post]
func RefreshHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "refresh token success"})
}
