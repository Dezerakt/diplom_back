package models

import (
	"gorm.io/gorm"
	"time"
)

type Singer struct {
	gorm.Model
	Name      string    `gorm:"type:string;unique;not null" json:"name"`
	BirthDate time.Time `json:"birth_date"`

	Albums []Album `json:"albums"`
}
type SingerInput struct {
	Name      string    `json:"name"`
	BirthDate time.Time `json:"birth_date"`

	Albums []Album `json:"albums,omitempty"`
}
