package model

import (
	"agahi-plus-plus/internal/postgres"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	UserID      uint           `json:"-"`
	User        *User          `json:"-"`
	Token       string         `gorm:"type:varchar(255); uniqueIndex" json:"token"`
	IsConnected bool           `gorm:"default: false" json:"is_connected"`
	Title       string         `gorm:"type:varchar(255)" json:"title"`
	Images      postgres.Jsonb `gorm:"type:jsonb" json:"images"`
	NewImages   postgres.Jsonb `gorm:"type:jsonb" json:"new_images"`
}
