package models

import "gorm.io/gorm"

type Album struct {
	gorm.Model
	SingerID    uint    `json:"singer_id" gorm:"type:uint"`
	SingerName  string  `gorm:"references:Name" json:"singer_name"`
	Name        string  `json:"name" gorm:"type:string"`
	Count       int     `json:"count" gorm:"type:int"`
	Price       int     `json:"price" gorm:"type:int"`
	ImageURL    string  `json:"image_url" gorm:"type:string"`
	ReleaseYear int     `json:"release_date" gorm:"type:int"`
	Genre       string  `json:"genre" gorm:"type:string"`
	Songs       []Song  `json:"songs"`
	Images      []Image `json:"images"`

	Cart Cart
}
