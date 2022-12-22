package controllers

import (
	"diplom_back/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AlbumController struct {
	DB *gorm.DB
}

func NewAlbumController(DB *gorm.DB) AlbumController {
	return AlbumController{DB}
}

func (obj *AlbumController) GetAll(context *gin.Context) {
	var albums []models.Album
	obj.DB.Find(&albums)

	context.JSON(200, gin.H{
		"content": &albums,
	})
}

func (obj *AlbumController) GetById(context *gin.Context) {
	var album models.Album

	obj.DB.First(&album, context.Param("id"))
	context.JSON(200, &album)
}
