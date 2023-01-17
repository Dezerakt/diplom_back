package routes

import (
	"diplom_back/controllers"
	"github.com/gin-gonic/gin"
)

type CartRouteController struct {
	cartController controllers.CartController
}

func NewCartRouteController(cartController controllers.CartController) CartRouteController {
	return CartRouteController{
		cartController: cartController,
	}
}

func (c CartRouteController) CartRoute(rg *gin.RouterGroup) {
	//router := rg.Group("cart")

}
