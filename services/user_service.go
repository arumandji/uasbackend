package services

import (
	"uas_backend/models"
	"uas_backend/repositories"
)

type UserService interface {
	CreateUser(u *models.User, plainPassword string) error
	GetUserByID(id string) (*models.User, error)
	ListUsers(offset, limit int) ([]models.User, error)
	UpdateUser(u *models.User) error
	DeleteUser(id uint) error
}

type userService struct {
	userRepo repositories.UserRepository
	authSvc  AuthService
}

func NewUserService(uRepo repositories.UserRepository, aSvc AuthService) UserService {
	return &userService{userRepo: uRepo, authSvc: aSvc}
}

func (s *userService) CreateUser(u *models.User, plainPassword string) error {
	hash, err := s.authSvc.HashPassword(plainPassword)
	if err != nil {
		return err
	}
	u.Password = hash
	return s.userRepo.Create(u)
}

func (s *userService) GetUserByID(id string) (*models.User, error) {
	return s.userRepo.FindByID(id)
}

func (s *userService) ListUsers(offset, limit int) ([]models.User, error) {
	return s.userRepo.List(offset, limit)
}

func (s *userService) UpdateUser(u *models.User) error {
	return s.userRepo.Update(u)
}

func (s *userService) DeleteUser(id uint) error {
	return s.userRepo.Delete(id)
}