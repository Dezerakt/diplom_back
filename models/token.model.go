package models

import "gorm.io/gorm"

type Token struct {
	gorm.Model
	UserID       uint   `json:"user_id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
