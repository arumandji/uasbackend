package models

import "time"

type Dosen struct {
	ID         string    `gorm:"type:uuid;primaryKey" json:"id"`
	UserID     string    `gorm:"type:uuid;not null;unique" json:"user_id"`
	DosenID string    `gorm:"size:20;unique;not null" json:"dosen_id"`
	Department string    `gorm:"size:100" json:"department"`
	CreatedAt  time.Time `json:"created_at"`
}
