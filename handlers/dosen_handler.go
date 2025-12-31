package handlers

import (
	"net/http"
	"uas_backend/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"uas_backend/services"
)

type DosenHandler struct {
	service services.DosenService
}

func NewDosenHandler(s services.DosenService) *DosenHandler {
	return &DosenHandler{service: s}
}

// CreateDosen godoc
// @Summary Create dosen
// @Description Membuat data dosen baru
// @Tags Dosen
// @Security BearerAuth
// @Accept multipart/form-data
// @Produce json
// @Param user_id formData string true "User ID"
// @Param nidn formData string true "NIDN Dosen"
// @Param nama_dosen formData string true "Nama Dosen"
// @Success 201 {object} models.Dosen
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/dosen [post]
func (h *DosenHandler) CreateDosen(c *gin.Context) {
    var m models.Dosen
    if err := c.ShouldBindJSON(&m); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // ID biarkan DB generate (integer)
    if err := h.service.Create(&m); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "dosen berhasil dibuat", "data": m})
}



// GetAllDosen godoc
// @Summary Get all dosen
// @Description Mengambil seluruh data dosen
// @Tags Dosen
// @Security BearerAuth
// @Produce json
// @Success 200 {array} models.Dosen
// @Failure 500 {object} map[string]string
// @Router /v1/dosen [get]
// GET /dosen
func (h *DosenHandler) GetAllDosen(c *gin.Context) {
	list, err := h.service.GetAllDosen()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch dosen",
		})
		return
	}

	c.JSON(http.StatusOK, list)
}

// GetDosenByID godoc
// @Summary Get detail dosen
// @Description Mengambil detail dosen berdasarkan ID
// @Tags Dosen
// @Security BearerAuth
// @Produce json
// @Param id path string true "Dosen ID"
// @Success 200 {object} models.Dosen
// @Failure 404 {object} map[string]string
// @Router /v1/dosen/{id} [get]
// GET /dosen/:id
func (h *DosenHandler) GetDosenByID(c *gin.Context) {
	id := c.Param("id")

	data, err := h.service.GetDosenByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "dosen not found",
		})
		return
	}

	c.JSON(http.StatusOK, data)
}

// GetDosenByUserID godoc
// @Summary Get dosen by user ID
// @Description Mengambil data dosen berdasarkan user ID
// @Tags Dosen
// @Security BearerAuth
// @Produce json
// @Param user_id path string true "User ID"
// @Success 200 {object} models.Dosen
// @Failure 404 {object} map[string]string
// @Router /v1/dosen/user/{user_id} [get]
// GET /dosen/user/:user_id
func (h *DosenHandler) GetDosenByUserID(c *gin.Context) {
	userID := c.Param("user_id")

	data, err := h.service.GetDosenByUserID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "dosen not found",
		})
		return
	}

	c.JSON(http.StatusOK, data)
}

// UpdateDosen godoc
// @Summary Update dosen
// @Description Update dosen information
// @Tags Dosen
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "Dosen ID"	
// @Param dosen body models.Dosen true "Dosen Data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/dosen/{id} [put]
func (h *DosenHandler) UpdateDosen(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	var m models.Dosen
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	m.ID = strconv.FormatUint(uint64(id), 10) // pastikan ID model bertipe string
	if err := h.service.Update(&m); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "dosen updated"})
}

// DeleteDosen godoc
// @Summary Delete dosen
// @Description Delete dosen by ID
// @Tags Dosen
// @Security BearerAuth
// @Param id path string true "Dosen ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/dosen/{id} [delete]
func (h *DosenHandler) DeleteDosen(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "dosen deleted"})
}