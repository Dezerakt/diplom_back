package models

import "gorm.io/gorm"

/*
TODO
ReleaseDate
Genre
SongList
SpotifyWidget
*/

type Album struct {
	gorm.Model
	SingerID uint
	Name     string `json:"name"`
	Count    int    `json:"count"`
	Price    int    `json:"price" gorm:"type:int"`
	ImageURL string `json:"image_url" gorm:"type:string"`
}
