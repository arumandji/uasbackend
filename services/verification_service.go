package services

import (
	"errors"
	"time"
	"uas_backend/repositories"
)

type VerificationService interface {
	VerifyAchievement(refID, verifierID string) error
	RejectAchievement(refID, verifierID, note string) error
}

type verificationService struct {
	refRepo repositories.AchievementRefRepository
}

func NewVerificationService(refRepo repositories.AchievementRefRepository) VerificationService {
	return &verificationService{refRepo: refRepo}
}

func (s *verificationService) VerifyAchievement(refID, verifierID string) error {
	ref, err := s.refRepo.FindByID(refID)
	if err != nil {
		return err
	}
	if ref.Status != "submitted" {
		return errors.New("only submitted can be verified")
	}
	now := time.Now()
	ref.Status = "verified"
	ref.VerifiedAt = &now
	ref.VerifiedBy = &verifierID
	return s.refRepo.Update(ref)
}

func (s *verificationService) RejectAchievement(refID, verifierID, note string) error {
	ref, err := s.refRepo.FindByID(refID)
	if err != nil {
		return err
	}
	if ref.Status != "submitted" {
		return errors.New("only submitted can be rejected")
	}
	now := time.Now()
	ref.Status = "rejected"
	ref.VerifiedAt = &now
	ref.VerifiedBy = &verifierID
	ref.RejectionNote = &note
	return s.refRepo.Update(ref)
}
