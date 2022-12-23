package controllers

import (
	"diplom_back/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

type SingerController struct {
	DB *gorm.DB
}

func NewSingerController(DB *gorm.DB) SingerController {
	return SingerController{DB}
}

func (obj SingerController) GetById(context *gin.Context) {
	var singers models.Singer
	err := obj.DB.Preload("Albums").First(&singers, context.Param("id")).Error
	if err != nil {
		log.Fatal(err)
	}
	context.JSON(200, &singers)
}
