package models

import (
	"gorm.io/gorm"
	"time"
)

type Singer struct {
	gorm.Model
	Name      string    `gorm:"type:string;unique;not null" json:"name"`
	BirthDate time.Time `gorm:"notnull"`

	Albums []Album `gorm:"constraint:OnDelete:CASCADE;ForeignKey:SingerID;"`
}
