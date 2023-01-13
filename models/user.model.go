package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"type:string" json:"login"`
	Username string `gorm:"type:string" json:"username"`
	Password string `gorm:"type:string" json:"password"`
	Token    Token
}

type SignInUser struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignUpUser struct {
	Email    string `gorm:"type:string" json:"email" binding:"required"`
	Username string `gorm:"type:string" json:"username" binding:"required"`
	Password string `gorm:"type:string" json:"password" binding:"required"`
}
