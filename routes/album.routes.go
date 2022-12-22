package routes

import (
	"diplom_back/controllers"
	"github.com/gin-gonic/gin"
)

type AlbumRouteController struct {
	albumController controllers.AlbumController
}

func NewAlbumRouteController(albumController controllers.AlbumController) AlbumRouteController {
	return AlbumRouteController{albumController}
}

func (rc *AlbumRouteController) AlbumRoute(rg *gin.RouterGroup) {
	router := rg.Group("/album")

	router.GET("/get-all", rc.albumController.GetAll)
	router.GET("/:id", rc.albumController.GetById)
}
