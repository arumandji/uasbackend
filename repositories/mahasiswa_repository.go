package repositories

import (
	"uas_backend/models"

	"gorm.io/gorm"
)

type MahasiswaRepository interface {
	FindByUserID(userID string) (*models.Mahasiswa, error)
	FindByID(id string) (*models.Mahasiswa, error)
	ListByAdvisor(dosenWaliId string) ([]models.Mahasiswa, error)
	CreateMahasiswa(s *models.Mahasiswa) error
	Update(s *models.Mahasiswa) error
	Delete(id uint) error
}


type mahasiswaRepo struct {
	db *gorm.DB
}

func NewMahasiswaRepository(db *gorm.DB) MahasiswaRepository {
	return &mahasiswaRepo{db: db}
}

func (r *mahasiswaRepo) FindByUserID(userID string) (*models.Mahasiswa, error) {
	var s models.Mahasiswa
	if err := r.db.Where("user_id = ?", userID).First(&s).Error; err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *mahasiswaRepo) FindByID(id string) (*models.Mahasiswa, error) {
	var s models.Mahasiswa
	if err := r.db.Where("id = ?", id).First(&s).Error; err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *mahasiswaRepo) ListByAdvisor(advisorID string) ([]models.Mahasiswa, error) {
	var list []models.Mahasiswa
	if err := r.db.Where("dosen_wali_id = ?", advisorID).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *mahasiswaRepo) CreateMahasiswa(s *models.Mahasiswa) error {
	return r.db.Create(s).Error
}

func (r *mahasiswaRepo) Update(s *models.Mahasiswa) error {
	return r.db.Save(s).Error
}

func (r *mahasiswaRepo) Delete(id uint) error {
	result := r.db.Delete(&models.Mahasiswa{}, id)
    return result.Error
}
