package services

import (
	"errors"
	"os"
	"time"
	"uas_backend/models"
	"uas_backend/repositories"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(username, password string) (string, *models.User, error)
	HashPassword(pw string) (string, error)
	CheckPassword(hash, pw string) error
}

type authService struct {
	userRepo repositories.UserRepository
	secret   string
	ttl      time.Duration
}

func NewAuthService(uRepo repositories.UserRepository) AuthService {
	ttl := time.Hour * 24
	secret := os.Getenv("JWT_SECRET")
	return &authService{userRepo: uRepo, secret: secret, ttl: ttl}
}

func (s *authService) HashPassword(pw string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	return string(bytes), err
}

func (s *authService) CheckPassword(hash, pw string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
}

func (s *authService) Login(username, password string) (string, *models.User, error) {
	u, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return "", nil, err
	}
	if !u.IsActive {
		return "", nil, errors.New("user not active")
	}
	if err := s.CheckPassword(u.Password, password); err != nil {
		return "", nil, errors.New("invalid credentials")
	}
	// Build token
	claims := jwt.MapClaims{
		"sub":  u.ID,
		"role": u.RoleID,
		"exp":  time.Now().Add(s.ttl).Unix(),
		"iat":  time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(s.secret))
	if err != nil {
		return "", nil, err
	}
	return ss, u, nil
}
