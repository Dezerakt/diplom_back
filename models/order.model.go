package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID  int   `json:"user_id"`
	User    User  `gorm:"constraint:OnDelete:CASCADE;"`
	AlbumID int   `json:"album_id"`
	Album   Album `gorm:"constraint:OnDelete:CASCADE;"`
}
