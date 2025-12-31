package services

import (
	"uas_backend/models"
	"uas_backend/repositories"
)

type MahasiswaService interface {
	GetMahasiswaByID(id string) (*models.Mahasiswa, error)
	GetMahasiswaByUserID(userID string) (*models.Mahasiswa, error)
	GetMahasiswaByAdvisor(advisorID string) ([]models.Mahasiswa, error)
	Create(m *models.Mahasiswa) error
	Update(m *models.Mahasiswa) error
	Delete(id uint) error
}

type mahasiswaService struct {
	repo repositories.MahasiswaRepository
}

func NewMahasiswaService(repo repositories.MahasiswaRepository) MahasiswaService {
	return &mahasiswaService{repo: repo}
}

func (s *mahasiswaService) GetMahasiswaByID(id string) (*models.Mahasiswa, error) {
	return s.repo.FindByID(id)
}

func (s *mahasiswaService) GetMahasiswaByUserID(userID string) (*models.Mahasiswa, error) {
	return s.repo.FindByUserID(userID)
}

func (s *mahasiswaService) GetMahasiswaByAdvisor(advisorID string) ([]models.Mahasiswa, error) {
	return s.repo.ListByAdvisor(advisorID)
}

func (s *mahasiswaService) Create(m *models.Mahasiswa) error {
	return s.repo.CreateMahasiswa(m)
}

func (s *mahasiswaService) Update(m *models.Mahasiswa) error {
	return s.repo.Update(m)
}

func (s *mahasiswaService) Delete(id uint) error {
	return s.repo.Delete(id)
}
