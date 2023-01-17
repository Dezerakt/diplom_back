package controllers

import (
	"gorm.io/gorm"
)

type CartController struct {
	DB *gorm.DB
}

func NewCartController(DB *gorm.DB) CartController {
	return CartController{DB: DB}
}
