package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Login    string `gorm:"type:string" json:"login"`
	Username string `gorm:"type:string" json:"username"`
	Password string `gorm:"type:string" json:"password"`
}

type SignInUser struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}
