package services

import (
	"uas_backend/models"
	"uas_backend/repositories"
)

type DosenService interface {
	GetAllDosen() ([]models.Dosen, error)
	GetDosenByID(id string) (*models.Dosen, error)
	GetDosenByUserID(userID string) (*models.Dosen, error)
	Create(d *models.Dosen) error
	Update(m *models.Dosen) error
	Delete(id uint) error
}

type dosenService struct {
	repo repositories.DosenRepository
}

func NewDosenService(repo repositories.DosenRepository) DosenService {
	return &dosenService{repo: repo}
}

func (s *dosenService) GetAllDosen() ([]models.Dosen, error) {
	return s.repo.ListAll()
}

func (s *dosenService) GetDosenByID(id string) (*models.Dosen, error) {
	return s.repo.FindByID(id)
}

func (s *dosenService) GetDosenByUserID(userID string) (*models.Dosen, error) {
	return s.repo.FindByUserID(userID)
}

func (s *dosenService) Update(m *models.Dosen) error {
	return s.repo.Update(m)
}

func (s *dosenService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *dosenService) Create(d *models.Dosen) error {
	return s.repo.Create(d)
}