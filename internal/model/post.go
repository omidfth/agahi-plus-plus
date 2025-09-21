package model

import (
	"agahi-plus-plus/internal/postgres"
	"encoding/json"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	UserID         uint           `json:"-"`
	User           *User          `json:"-"`
	Token          string         `gorm:"type:varchar(255); uniqueIndex" json:"token"`
	IsConnected    bool           `gorm:"default: false" json:"is_connected"`
	Title          string         `gorm:"type:varchar(255)" json:"title"`
	Description    string         `gorm:"type:text" json:"description"`
	Images         postgres.Jsonb `gorm:"type:jsonb" json:"images"`
	NewImages      postgres.Jsonb `gorm:"type:jsonb" json:"new_images"`
	SelectedImages postgres.Jsonb `gorm:"type:jsonb" json:"selected_images"`
}

func (m *Post) GetImages() []string {
	j, _ := json.Marshal(m.Images)
	var val []string
	err := json.Unmarshal(j, &val)
	if err != nil {
		return nil
	}

	return val
}

func (m *Post) GetNewImages() []string {
	j, _ := json.Marshal(m.NewImages)
	var val []string
	err := json.Unmarshal(j, &val)
	if err != nil {
		return nil
	}

	return val
}

func (m *Post) GetSelectedImages() []string {
	j, _ := json.Marshal(m.SelectedImages)
	var val []string
	err := json.Unmarshal(j, &val)
	if err != nil {
		return nil
	}

	return val
}

func (m *Post) GetAllImages() []string {
	images := m.GetImages()
	newImages := m.GetNewImages()
	selectedImages := m.GetSelectedImages()

	var ret []string

	for _, image := range images {
		if m.contains(image, selectedImages) {
			continue
		}

		ret = append(ret, image)
	}

	ret = append(ret, newImages...)

	return ret
}

func (m *Post) contains(s string, ss []string) bool {
	for _, s1 := range ss {
		if s == s1 {
			return true
		}
	}
	return false
}
