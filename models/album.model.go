package models

import "gorm.io/gorm"

type Album struct {
	gorm.Model
	SingerID int    `json:"singer_id"`
	Price    int    `json:"price" gorm:"type:int"`
	ImageURL string `json:"image_url" gorm:"type:string"`
}
