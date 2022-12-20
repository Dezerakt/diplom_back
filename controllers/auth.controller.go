package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController struct {
	DB *gorm.DB
}

func NewAuthController(DB *gorm.DB) AuthController {
	return AuthController{DB}
}

func (c *AuthController) SignInUser(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"response": "success",
	})
}
