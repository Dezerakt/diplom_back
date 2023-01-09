package controllers

import (
	"diplom_back/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

type AlbumController struct {
	DB *gorm.DB
}

func NewAlbumController(DB *gorm.DB) AlbumController {
	return AlbumController{DB}
}

func (obj *AlbumController) GetAll(context *gin.Context) {
	var albums []models.Album
	obj.DB.Preload("Images").Preload("Songs").Find(&albums)

	context.JSON(200, &albums)
}

func (obj *AlbumController) GetById(context *gin.Context) {
	var album models.Album

	obj.DB.First(&album, context.Param("id"))
	context.JSON(200, &album)
}

func (obj *AlbumController) AddAlbum(context *gin.Context) {
	var album models.Album

	if err := context.ShouldBindJSON(&album); err != nil {
		context.IndentedJSON(400, err)
		log.Fatal(err)
	}

	err := obj.DB.Create(&models.Album{
		SingerID: album.SingerID,
		Name:     album.Name,
		Count:    album.Count,
		Price:    album.Price,
		ImageURL: album.ImageURL,
	}).Error

	if err != nil {
		context.IndentedJSON(400, err)
		log.Fatal(err)
	}

	context.IndentedJSON(200, album)
}

func (obj *AlbumController) UpdateAlbum(context *gin.Context) {
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

func (obj *AlbumController) DeleteAlbum(context *gin.Context) {
	err := obj.DB.Unscoped().Delete(&models.Album{}, context.Param("id")).Error

	if err != nil {
		log.Fatal(err)
	}

	context.JSON(200, gin.H{
		"status": "success",
	})
}
