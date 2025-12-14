package services

import (
	"errors"
	"time"
	"uas_backend/models"
	"uas_backend/repositories"

)

type AchievementService interface {
	CreateAchievement(a *models.Achievement) (string, error)
	GetAchievementByID(mongoID string) (*models.Achievement, error)
	SubmitForVerification(refID string) error
}

type achievementService struct {
	achRepo      repositories.AchievementMongoRepository
	refRepo      repositories.AchievementRefRepository
	mahasiswaRepo  repositories.MahasiswaRepository
}

func NewAchievementService(achRepo repositories.AchievementMongoRepository, refRepo repositories.AchievementRefRepository, stuRepo repositories.MahasiswaRepository) AchievementService {
	return &achievementService{achRepo: achRepo, refRepo: refRepo, mahasiswaRepo: stuRepo}
}

func (s *achievementService) CreateAchievement(a *models.Achievement) (string, error) {
	if a.MahasiswaID == "" {
		return "", errors.New("mahasiswa_id required")
	}
	id, err := s.achRepo.Create(a)
	if err != nil {
		return "", err
	}
	// Create reference in Postgres
	ref := &models.AchievementReference{
		ID:               generateUUID(),
		MahasiswaID:        a.MahasiswaID,
		MongoAchievement: id,
		Status:           "draft",
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}
	if err := s.refRepo.Create(ref); err != nil {
		// optional: rollback mongo doc (not implemented)
		return "", err
	}
	return id, nil
}

func (s *achievementService) GetAchievementByID(mongoID string) (*models.Achievement, error) {
	return s.achRepo.FindByID(mongoID)
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
	return s.refRepo.Update(ref)
}

// generateUUID is a simple generator wrapper, you may replace with github.com/google/uuid
func generateUUID() string {
	return time.Now().Format("20060102150405") // placeholder; replace with uuid.NewString()
}
