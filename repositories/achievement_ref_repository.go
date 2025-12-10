package repositories

import (
	"uas_backend/models"
	"time"

	"gorm.io/gorm"
)

type AchievementRefRepository interface {
	Create(ref *models.AchievementReference) error
	FindByID(id string) (*models.AchievementReference, error)
	ListByStudentIDs(studentIDs []string, offset, limit int) ([]models.AchievementReference, error)
	Update(ref *models.AchievementReference) error
}

type achRefRepo struct {
	db *gorm.DB
}

func NewAchievementRefRepository(db *gorm.DB) AchievementRefRepository {
	return &achRefRepo{db: db}
}

func (r *achRefRepo) Create(ref *models.AchievementReference) error {
	now := time.Now()
	ref.CreatedAt = now
	ref.UpdatedAt = now
	return r.db.Create(ref).Error
}

func (r *achRefRepo) FindByID(id string) (*models.AchievementReference, error) {
	var ref models.AchievementReference
	if err := r.db.Where("id = ?", id).First(&ref).Error; err != nil {
		return nil, err
	}
	return &ref, nil
}

func (r *achRefRepo) ListByStudentIDs(studentIDs []string, offset, limit int) ([]models.AchievementReference, error) {
	var list []models.AchievementReference
	if err := r.db.Where("student_id IN ?", studentIDs).Offset(offset).Limit(limit).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *achRefRepo) Update(ref *models.AchievementReference) error {
	ref.UpdatedAt = time.Now()
	return r.db.Save(ref).Error
}
