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

func (obj *SingerController) GetById(context *gin.Context) {
	var singers models.Singer
	err := obj.DB.Preload("Albums").First(&singers, context.Param("id")).Error
	if err != nil {
		log.Fatal(err)
	}
	context.JSON(200, &singers)
}

func (obj *SingerController) AddSinger(context *gin.Context) {
	var singer models.Singer

	if err := context.ShouldBindJSON(&singer); err != nil {
		context.IndentedJSON(400, err)
		log.Fatal(err)
	}

	err := obj.DB.Create(&models.Singer{
		Model:     gorm.Model{},
		Name:      singer.Name,
		BirthDate: singer.BirthDate,
		Albums:    singer.Albums,
	}).Error

	if err != nil {
		context.IndentedJSON(400, err)
		log.Fatal(err)
	}

	context.IndentedJSON(200, &singer)
}

func (obj *SingerController) UpdateAlbum(context *gin.Context) {
	var updateAlbum models.Album

	if err := context.ShouldBindJSON(&updateAlbum); err != nil {
		log.Fatal(err)
	}

	if err := obj.DB.Model(&models.Album{}).Where("id = ?", context.Param("id")).Updates(&models.Album{
		SingerID: updateAlbum.SingerID,
		Name:     updateAlbum.Name,
		Count:    updateAlbum.Count,
		Price:    updateAlbum.Price,
		ImageURL: updateAlbum.ImageURL,
	}).Error; err != nil {
		log.Fatal("Update failed")
	}

	context.JSON(200, "Successfully")
}

func (obj *SingerController) DeleteAlbum(context *gin.Context) {
	err := obj.DB.Unscoped().Delete(&models.Album{}, context.Param("id")).Error

	if err != nil {
		log.Fatal(err)
	}

	context.JSON(200, gin.H{
		"status": "success",
	})
}
