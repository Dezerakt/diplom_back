package models

import "gorm.io/gorm"

type Album struct {
	gorm.Model
	SingerID int    `json:"singer_id"`
	Name     string `json:"name"`
	Count    int    `json:"count"`
	Price    int    `json:"price" gorm:"type:int"`
	ImageURL string `json:"image_url" gorm:"type:string"`
}
