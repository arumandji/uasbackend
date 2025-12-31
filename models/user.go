package models

import (

)

type User struct {
	ID           string         `gorm:"type:uuid;primaryKey" json:"id"`
	Username     string         `gorm:"size:50;unique;not null" json:"username"`
	Password	 string         `gorm:"size:255;not null" json:"password"`
	Nama		 string         `gorm:"size:100;not null" json:"nama"`
	RoleID       string         `gorm:"type:uuid" json:"role_id"`
	IsActive     bool           `gorm:"default:true" json:"is_active"`
}