package models

import "gorm.io/gorm"

type Song struct {
	gorm.Model
	AlbumID       uint   `json:"album_id"`
	Name          string `json:"name"`
	SpotifyWidget string `json:"spotify_widget"`
}
