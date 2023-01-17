package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	AlbumID int `json:"album_id"`
	UserID  int `json:"user_id"`
}
