package models

import "gorm.io/gorm"

type Song struct {
	gorm.Model
	AlbumID       uint   `json:"album_id"`
	Duration      string `json:"duration"`
	Name          string `json:"name"`
	SpotifyWidget string `json:"spotify_widget"`
}
