package models

import "time"

type Student struct {
	ID           string    `gorm:"type:uuid;primaryKey" json:"id"`
	UserID       string    `gorm:"type:uuid;not null;unique" json:"user_id"`
	MahasiswaID    string    `gorm:"size:20;unique;not null" json:"mahasiswa_id"`
	ProgramStudy string    `gorm:"size:100" json:"program_study"`
	AcademicYear string    `gorm:"size:10" json:"academic_year"`
	AdvisorID    string    `gorm:"type:uuid" json:"advisor_id"` // lecturer id
	CreatedAt    time.Time `json:"created_at"`
}