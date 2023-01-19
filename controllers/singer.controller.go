package controllers

import (
	"diplom_back/models"
	"diplom_back/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
		utils.ErrorResponse(context, err)
	}
	context.JSON(200, &singers)
}

func (obj *SingerController) AddSinger(context *gin.Context) {
	var singer models.Singer

	if err := context.ShouldBindJSON(&singer); err != nil {
		utils.ErrorResponse(context, err)
	}

	err := obj.DB.Create(&models.Singer{
		Model:     gorm.Model{},
		Name:      singer.Name,
		BirthDate: singer.BirthDate,
		Albums:    singer.Albums,
	}).Error

	if err != nil {
		utils.ErrorResponse(context, err)
	}

	context.IndentedJSON(200, &singer)
}

func (obj *SingerController) UpdateAlbum(context *gin.Context) {
	var updateAlbum models.Album

	if err := context.ShouldBindJSON(&updateAlbum); err != nil {
		utils.ErrorResponse(context, err)
	}

	if err := obj.DB.Model(&models.Album{}).Where("id = ?", context.Param("id")).Updates(&models.Album{
		SingerID: updateAlbum.SingerID,
		Name:     updateAlbum.Name,
		Count:    updateAlbum.Count,
		Price:    updateAlbum.Price,
		ImageURL: updateAlbum.ImageURL,
	}).Error; err != nil {
		utils.ErrorResponse(context, err)
	}

	context.JSON(200, "Successfully")
}

func (obj *SingerController) DeleteAlbum(context *gin.Context) {
	err := obj.DB.Unscoped().Delete(&models.Album{}, context.Param("id")).Error

	if err != nil {
		utils.ErrorResponse(context, err)
	}

	context.JSON(200, gin.H{
		"status": "success",
	})
}
