package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID  int `json:"user_id"`
	User    User
	AlbumID int `json:"album_id"`
	Album   Album
}
