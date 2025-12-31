package repositories

import (
	"uas_backend/models"

	"gorm.io/gorm"
)

type DosenRepository interface {
	FindByUserID(userID string) (*models.Dosen, error)
	FindByID(id string) (*models.Dosen, error)
	ListAll() ([]models.Dosen, error)
	Create(d *models.Dosen) error
	Update(d *models.Dosen) error
	Delete(id uint) error
}


type dosenRepo struct {
	db *gorm.DB
}

func NewDosenRepository(db *gorm.DB) DosenRepository {
	return &dosenRepo{db: db}
}

func (r *dosenRepo) FindByUserID(userID string) (*models.Dosen, error) {
	var l models.Dosen
	if err := r.db.Where("user_id = ?", userID).First(&l).Error; err != nil {
		return nil, err
	}
	return &l, nil
}

func (r *dosenRepo) FindByID(id string) (*models.Dosen, error) {
	var l models.Dosen
	if err := r.db.Where("id = ?", id).First(&l).Error; err != nil {
		return nil, err
	}
	return &l, nil
}

func (r *dosenRepo) ListAll() ([]models.Dosen, error) {
	var list []models.Dosen
	if err := r.db.Find(&list).Error; err != nil {
		return list, err
	}
	return list, nil
}

func (r *dosenRepo) Create(l *models.Dosen) error {
	return r.db.Create(l).Error
}

func (r *dosenRepo) Update(l *models.Dosen) error {
	return r.db.Save(l).Error
}

func (r *dosenRepo) Delete(id uint) error {
	result := r.db.Delete(&models.Dosen{}, id)
    return result.Error
}
