package controllers

import (
	"diplom_back/models"
	"diplom_back/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"strconv"
	"time"
)

type CartController struct {
	DB *gorm.DB
}

func NewCartController(DB *gorm.DB) CartController {
	return CartController{DB: DB}
}

func (receiver *CartController) AddFew(ctx *gin.Context) {
	type Albums struct {
		AlbumsId []string `json:"album_Ids"`
		UserId   int      `json:"user_id"`
	}

	var albumIds Albums

	if err := ctx.ShouldBindJSON(&albumIds); err != nil {
		utils.ErrorResponse(ctx, err)
	}

	for _, value := range albumIds.AlbumsId {
		if value != "" {
			intValue, err := strconv.Atoi(value)

			if err != nil {
				utils.ErrorResponse(ctx, err)
			}

			err = receiver.DB.Create(&models.Cart{
				Model:       gorm.Model{},
				AlbumID:     intValue,
				UserID:      albumIds.UserId,
				ReleaseDate: time.Now().Format("02/01/2006"),
			}).Error

			if err != nil {
				utils.ErrorResponse(ctx, err)
			}
		}
	}

	ctx.JSON(200, gin.H{
		"status": "success",
	})
}

func (receiver *CartController) GetAllByUser(ctx *gin.Context) {
	type UserId struct {
		UserId int `json:"user_id"`
	}

	type OrderAlbum struct {
		Cart  models.Cart
		Album models.Album
	}

	var orderAlbums []OrderAlbum
	var orders []models.Cart
	var userId UserId

	if err := ctx.ShouldBindJSON(&userId); err != nil {
		utils.ErrorResponse(ctx, err)
		log.Fatal(err.Error())
	}

	receiver.DB.Find(&orders, &models.Cart{UserID: userId.UserId})

	for _, order := range orders {
		var tempAlbum models.Album
		receiver.DB.First(&tempAlbum, &models.Album{
			Model: gorm.Model{ID: uint(order.AlbumID)},
		})
		orderAlbums = append(orderAlbums, OrderAlbum{Cart: order, Album: tempAlbum})
	}

	ctx.JSON(200, &orderAlbums)
}
