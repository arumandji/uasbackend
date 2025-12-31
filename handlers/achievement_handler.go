package handlers

import (
	"net/http"
	"strconv"
	
	"uas_backend/database"
	"uas_backend/models"
	"github.com/gin-gonic/gin"
)


// CreateAchievement godoc
// @Summary Create achievement
// @Tags Achievement
// @Security BearerAuth
// @Accept multipart/form-data
// @Produce json
// @Param judul formData string true "Judul Achievement"
// @Param tingkat formData string true "Tingkat Achievement"
// @Param kategori formData string true "Kategori Achievement"
// @Param tahun formData int true "Tahun Achievement"
// @Param keterangan formData string false "Keterangan / Alasan Reject"
// @Success 201 {object} models.Achievement
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/achievements [post]
func CreateAchievement(c *gin.Context) {
	judul := c.PostForm("judul")
	tingkat := c.PostForm("tingkat")
	kategori := c.PostForm("kategori")
	tahunStr := c.PostForm("tahun")
	keterangan := c.PostForm("keterangan")

	if judul == "" || tingkat == "" || kategori == "" || tahunStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "judul, tingkat, kategori, dan tahun wajib diisi"})
		return
	}

	tahun, err := strconv.Atoi(tahunStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "tahun harus angka"})
		return
	}

	// Ambil mahasiswa_id dari token
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user_id not found in token"})
		return
	}

	achievement := models.Achievement{
		MahasiswaID: userID.(string),
		Judul:       judul,
		Tingkat:     tingkat,
		Kategori:    kategori,
		Tahun:       tahun,
		Status:      "submitted",
		Keterangan:  keterangan,
	}

	if err := database.DB.Create(&achievement).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create achievement: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, achievement)
}

// GetAllAchievements godoc
// @Summary Get all achievements
// @Tags Achievement
// @Security BearerAuth
// @Produce json
// @Success 200 {array} models.Achievement
// @Router /v1/achievements [get]
// GetAllAchievements handles GET /achievements
func GetAllAchievements(c *gin.Context) {
	var achievements []models.Achievement

	if err := database.DB.Find(&achievements).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch achievements",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": achievements,
	})
}

// GetMyAchievements godoc
// @Summary Get my achievements
// @Tags Achievement
// @Security BearerAuth
// @Produce json
// @Success 200 {array} models.Achievement
// @Router /v1/achievements [get]
// GetMyAchievements handles GET /achievements for the current user
func GetMyAchievements(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var achievements []models.Achievement
	if err := database.DB.
		Where("mahasiswa_id = ?", userID).
		Find(&achievements).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch achievements",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": achievements,
	})
}

// UpdateAchievement godoc
// @Summary Update achievement
// @Description Update achievement by ID
// @Tags Achievement
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "Achievement ID"
// @Param achievement body models.Achievement true "Achievement Data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/achievements/{id} [put]
func UpdateAchievement(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	var req models.Achievement
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update record by ID
	if err := database.DB.Model(&models.Achievement{}).
		Where("id = ?", id).
		Updates(req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update achievement"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "achievement updated successfully"})
}

// DeleteAchievement godoc
// @Summary Delete achievement
// @Description Delete achievement by ID
// @Tags Achievement
// @Security BearerAuth
// @Produce json
// @Param id path string true "Achievement ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/achievements/{id} [delete]
func DeleteAchievement(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	if err := database.DB.Delete(&models.Achievement{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete achievement"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "achievement deleted successfully"})
}