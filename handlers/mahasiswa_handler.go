package handlers

import (
	"net/http"
	"strconv"
	"uas_backend/database"
	"uas_backend/models"

	"uas_backend/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MahasiswaHandler struct {
	service services.MahasiswaService
}

func NewMahasiswaHandler(s services.MahasiswaService) *MahasiswaHandler {
	return &MahasiswaHandler{service: s}
}

// GetMahasiswaByID godoc
// @Summary Get detail mahasiswa
// @Description Mengambil detail mahasiswa berdasarkan ID
// @Tags Mahasiswa
// @Security BearerAuth
// @Produce json
// @Param id path string true "Mahasiswa ID"
// @Success 200 {object} models.Mahasiswa
// @Failure 404 {object} map[string]string
// @Router /v1/mahasiswa/{id} [get]
// GET /mahasiswa/:id
func (h *MahasiswaHandler) GetMahasiswaByID(c *gin.Context) {
	id := c.Param("id")

	data, err := h.service.GetMahasiswaByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "mahasiswa not found",
		})
		return
	}

	c.JSON(http.StatusOK, data)
}

// GetMahasiswaByUserID godoc
// @Summary Get mahasiswa by user ID
// @Description Mengambil data mahasiswa berdasarkan user ID
// @Tags Mahasiswa
// @Security BearerAuth
// @Produce json
// @Param user_id path string true "User ID"
// @Success 200 {object} models.Mahasiswa
// @Failure 404 {object} map[string]string
// @Router /v1/mahasiswa/user/{user_id} [get]
// GET /mahasiswa/user/:user_id
func (h *MahasiswaHandler) GetMahasiswaByUserID(c *gin.Context) {
	userID := c.Param("user_id")

	data, err := h.service.GetMahasiswaByUserID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "mahasiswa not found",
		})
		return
	}

	c.JSON(http.StatusOK, data)
}

// GetMahasiswaByAdvisor godoc
// @Summary Get mahasiswa bimbingan dosen
// @Description Mengambil daftar mahasiswa berdasarkan dosen wali
// @Tags Mahasiswa
// @Security BearerAuth
// @Produce json
// @Param Dosen_wali_id path string true "Dosen Wali ID"
// @Success 200 {array} models.Mahasiswa
// @Failure 500 {object} map[string]string
// @Router /v1/mahasiswa/advisor/{dosen_wali_id} [get]
// GET /mahasiswa/advisor/:dosen_wali_id
func (h *MahasiswaHandler) GetMahasiswaByAdvisor(c *gin.Context) {
	advisorID := c.Param("dosen_wali_id")

	list, err := h.service.GetMahasiswaByAdvisor(advisorID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch mahasiswa",
		})
		return
	}

	c.JSON(http.StatusOK, list)
}

// CreateMahasiswa godoc
// @Summary Create mahasiswa
// @Description Membuat data mahasiswa baru
// @Tags Mahasiswa
// @Security BearerAuth
// @Accept multipart/form-data
// @Produce json
// @Param user_id formData string true "User ID"
// @Param nim formData string true "NIM Mahasiswa"
// @Param nama_mhs formData string true "Nama Mahasiswa"
// @Param angkatan formData string true "Angkatan"
// @Param prodi formData string true "Program Studi"
// @Param dosen_wali_id formData string false "ID Dosen Wali"
// @Success 201 {object} models.Mahasiswa
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/mahasiswa [post]
func (h *MahasiswaHandler) CreateMahasiswa(c *gin.Context) {
	namaMhs := c.PostForm("nama_mhs")
	nim := c.PostForm("nim")
	angkatan := c.PostForm("angkatan")
	prodi := c.PostForm("prodi")
	dosenIDStr := c.PostForm("dosen_wali_id")
	var dosenID uint
	if dosenIDStr != "" {
		parsedID, err := strconv.ParseUint(dosenIDStr, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid dosen_wali_id"})
			return
		}
		dosenID = uint(parsedID)
	}

	if namaMhs == "" || nim == "" || angkatan == "" || prodi == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "nama_mhs, nim, angkatan, prodi wajib diisi"})
		return
	}

	userID, _ := c.Get("user_id")

	m := models.Mahasiswa{
		ID:          uuid.New().String(), // hanya kalau DB tipe uuid
		UserID:      userID.(string),
		NamaMhs:     namaMhs,
		NIM:         nim,
		Angkatan:    angkatan,
		Prodi:       prodi,
		DosenWaliID: dosenID,
	}

	if err := database.DB.Create(&m).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "mahasiswa berhasil dibuat", "data": m})
}

// UpdateMahasiswa godoc
// @Summary Update mahasiswa
// @Description Mengupdate data mahasiswa berdasarkan ID
// @Tags Mahasiswa
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "Mahasiswa ID"
// @Param mahasiswa body models.Mahasiswa true "Mahasiswa Data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/mahasiswa/{id} [put]
func (h *MahasiswaHandler) UpdateMahasiswa(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	var m models.Mahasiswa
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	m.ID = strconv.Itoa(id)
	if err := h.service.Update(&m); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "mahasiswa updated"})
}

// DeleteMahasiswa godoc
// @Summary Delete mahasiswa
// @Description Menghapus data mahasiswa berdasarkan ID
// @Tags Mahasiswa
// @Security BearerAuth
// @Produce json
// @Param id path string true "Mahasiswa ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/mahasiswa/{id} [delete]
func (h *MahasiswaHandler) DeleteMahasiswa(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "mahasiswa deleted"})
}
