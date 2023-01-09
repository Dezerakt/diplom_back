package models

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	AlbumID uint   `json:"album_id" gorm:"type:uint"`
	URL     string `json:"url" gorm:"type:string"`
}
