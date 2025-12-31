package models

import (
)

type Dosen struct {
	ID         string    `gorm:"primaryKey" json:"id"`
	UserID     string    `gorm:"not null;unique" json:"user_id"`
	NIDN       string 	`gorm:"column:nidn;size:20;not null" json:"nidn"`
	NamaDosen  string 	`gorm:"column:nama_dosen;size:100" json:"nama_dosen"`
}

func (Dosen) TableName() string {
	return "dosen"
}