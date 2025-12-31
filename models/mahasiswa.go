package models

import (
)

type Mahasiswa struct {
	ID           string    	`gorm:"type:uuid;primaryKey" json:"id"`
	UserID       string    	`gorm:"not null;unique" json:"user_id"`
	NIM          string 	`gorm:"column:nim;size:20;not null" json:"nim"`
	NamaMhs      string 	`gorm:"column:nama_mhs;size:100" json:"nama_mhs"`
	Angkatan     string    	`gorm:"column:angkatan" json:"angkatan"`
	DosenWaliID  uint   	`gorm:"column:dosen_wali_id" json:"dosen_wali_id"`
	Prodi        string 	`gorm:"column:prodi;size:100" json:"prodi"`
}

func (Mahasiswa) TableName() string {
	return "mahasiswa"
}