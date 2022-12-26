package routes

import (
	"diplom_back/controllers"
	"github.com/gin-gonic/gin"
)

type SingerRouteController struct {
	singerController controllers.SingerController
}

func NewSingerRouteController(singerController controllers.SingerController) SingerRouteController {
	return SingerRouteController{
		singerController: singerController,
	}
}

func (rc *SingerRouteController) SingerRoute(rg *gin.RouterGroup) {
	router := rg.Group("singer")

	router.GET("/:id", rc.singerController.GetById)
	router.POST("/add", rc.singerController.AddSinger)
	router.PUT("/:id", rc.singerController.UpdateAlbum)
	router.DELETE("/:id", rc.singerController.DeleteAlbum)
}
