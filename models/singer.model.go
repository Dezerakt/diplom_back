package models

import (
	"gorm.io/gorm"
	"time"
)

type Singer struct {
	gorm.Model
	Name        string    `gorm:"type:string;unique;not null" json:"name"`
	BirthDate   time.Time `json:"birth_date"`
	ImageURL    string    `json:"image_url"`
	Description string    `json:"description"`

	Albums []Album `json:"albums"`
}
