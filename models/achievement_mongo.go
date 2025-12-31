package models

import "time"

func (Achievement) TableName() string {
	return "prestasi" // atau "prestasi"
}

// Mongo model for achievements (dynamic details as map[string]interface{})
type Achievement struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	MahasiswaID string   `gorm:"column:mahasiswa_id" json:"mahasiswa_id"`
	Judul       string `gorm:"column:judul" json:"judul"`
	Tingkat     string `gorm:"column:tingkat" json:"tingkat"`
	Kategori    string `gorm:"column:kategori" json:"kategori"`
	Tahun       int    `gorm:"column:tahun" json:"tahun"`
	Status      string `gorm:"column:status" json:"status"`       // submitted / verified / rejected
	Keterangan  string `gorm:"column:keterangan" json:"keterangan"` // alasan reject
}

type Attachment struct {
	FileName   string    `bson:"fileName" json:"file_name"`
	FileURL    string    `bson:"fileUrl" json:"file_url"`
	FileType   string    `bson:"fileType" json:"file_type"`
	UploadedAt time.Time `bson:"uploadedAt" json:"uploaded_at"`
}
