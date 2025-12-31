package services

import (
	"context"
	"errors"
	"time"

	"uas_backend/models"
	"uas_backend/repositories"
)

type AchievementService interface {
	CreateAchievement(a *models.Achievement) error
	GetAchievementByID(mongoID string) (map[string]interface{}, error)
	SubmitForVerification(refID string) error
	Update(a *models.Achievement) error
	Delete(id uint) error
}

type achievementService struct {
	achRepo        *repositories.AchievementMongoRepository
	refRepo        repositories.AchievementRefRepository
	mahasiswaRepo  repositories.MahasiswaRepository
}

func NewAchievementService(
	achRepo *repositories.AchievementMongoRepository,
	refRepo repositories.AchievementRefRepository,
	stuRepo repositories.MahasiswaRepository,
) AchievementService {
	return &achievementService{
		achRepo: achRepo,
		refRepo: refRepo,
		mahasiswaRepo: stuRepo,
	}
}

func (s *achievementService) CreateAchievement(a *models.Achievement) error {
	if a.MahasiswaID == "" {
		return errors.New("mahasiswa_id required")
	}

	// simpan ke MongoDB
	if err := s.achRepo.Create(context.Background(), a); err != nil {
		return err
	}

	// simpan referensi di PostgreSQL
	ref := &models.AchievementReference{
		ID:          generateUUID(),
		MahasiswaID: a.MahasiswaID,
		Status:      "draft",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	return s.refRepo.Create(ref)
}

func (s *achievementService) GetAchievementByID(mongoID string) (map[string]interface{}, error) {
	return s.achRepo.FindByID(context.Background(), mongoID)
}

func (s *achievementService) SubmitForVerification(refID string) error {
	ref, err := s.refRepo.FindByID(refID)
	if err != nil {
		return err
	}

	if ref.Status != "draft" {
		return errors.New("only draft can be submitted")
	}

	now := time.Now()
	ref.Status = "submitted"
	ref.SubmittedAt = &now
	ref.UpdatedAt = now

	return s.refRepo.Update(ref)
}

// sederhana & cukup untuk UAS
func generateUUID() string {
	return time.Now().Format("20060102150405")
}

func (s *achievementService) Update(a *models.Achievement) error {
	return s.achRepo.Update(context.Background(), a)
}

func (s *achievementService) Delete(id uint) error {
	return s.achRepo.Delete(context.Background(), id)
}