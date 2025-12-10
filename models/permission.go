package models

type Permission struct {
	ID          string `gorm:"type:uuid;primaryKey" json:"id"`
	Name        string `gorm:"size:100;unique;not null" json:"name"`
	Resource    string `gorm:"size:50" json:"resource"`
	Action      string `gorm:"size:50" json:"action"`
	Description string `json:"description"`
}