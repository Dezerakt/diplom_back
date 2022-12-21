package models

import "gorm.io/gorm"

type Singer struct {
	gorm.Model
	Name string `gorm:"type:string" json:"name"`

	Albums []Album
}
