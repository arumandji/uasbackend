package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           string         `gorm:"type:uuid;primaryKey" json:"id"`
	Username     string         `gorm:"size:50;unique;not null" json:"username"`
	Email        string         `gorm:"size:100;unique;not null" json:"email"`
	PasswordHash string         `gorm:"size:255;not null" json:"-"`
	FullName     string         `gorm:"size:100;not null" json:"full_name"`
	RoleID       string         `gorm:"type:uuid" json:"role_id"`
	IsActive     bool           `gorm:"default:true" json:"is_active"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}