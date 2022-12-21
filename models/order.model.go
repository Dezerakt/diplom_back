package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID  int
	User    User
	AlbumID int
	Album   Album
}
