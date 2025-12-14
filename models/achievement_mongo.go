package models

import "time"

// Mongo model for achievements (dynamic details as map[string]interface{})
type Achievement struct {
	ID              string                 `bson:"_id,omitempty" json:"id"`
	MahasiswaID     string                 `bson:"mahasiswaId" json:"mahasiswa_id"`
	AchievementType string                 `bson:"achievementType" json:"achievement_type"`
	Title           string                 `bson:"title" json:"title"`
	Description     string                 `bson:"description" json:"description"`
	Details         map[string]interface{} `bson:"details,omitempty" json:"details"`
	Attachments     []Attachment           `bson:"attachments,omitempty" json:"attachments"`
	Tags            []string               `bson:"tags,omitempty" json:"tags"`
	Points          *int                   `bson:"points,omitempty" json:"points"`
	CreatedAt       time.Time              `bson:"createdAt" json:"created_at"`
	UpdatedAt       time.Time              `bson:"updatedAt" json:"updated_at"`
}

type Attachment struct {
	FileName   string    `bson:"fileName" json:"file_name"`
	FileURL    string    `bson:"fileUrl" json:"file_url"`
	FileType   string    `bson:"fileType" json:"file_type"`
	UploadedAt time.Time `bson:"uploadedAt" json:"uploaded_at"`
}
