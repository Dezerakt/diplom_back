package models

import (
	"gorm.io/gorm"
)

type Token struct {
	gorm.Model
	UserID uint   `json:"user_id" gorm:"unique"`
	Token  string `json:"token" gorm:"type:string;"`
}
