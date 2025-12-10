package models

import "time"

// Reference to the document stored in MongoDB
type AchievementReference struct {
	ID               string    `gorm:"type:uuid;primaryKey" json:"id"`
	MahasiswaID        string    `gorm:"type:uuid;not null" json:"mahasiswa_id"`
	MongoAchievement string    `gorm:"size:50;not null" json:"mongo_achievement_id"`
	Status           string    `gorm:"type:VARCHAR(20);default:'draft'" json:"status"` // draft, submitted, verified, rejected
	SubmittedAt      *time.Time `json:"submitted_at"`
	VerifiedAt       *time.Time `json:"verified_at"`
	VerifiedBy       *string    `gorm:"type:uuid" json:"verified_by"`
	RejectionNote    *string    `json:"rejection_note"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
}
