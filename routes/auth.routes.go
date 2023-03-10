package routes

import (
	"diplom_back/controllers"
	"github.com/gin-gonic/gin"
)

type AuthRouteController struct {
	authController controllers.AuthController
}

func NewAuthRouteController(authController controllers.AuthController) AuthRouteController {
	return AuthRouteController{authController}
}

func (rc *AuthRouteController) AuthRoute(rg *gin.RouterGroup) {
	router := rg.Group("/auth")

	router.POST("/sign-in", rc.authController.SignInUser)
	router.POST("/sign-up", rc.authController.SignUpUser)
	router.POST("/logout", rc.authController.Logout)
	router.POST("/change", rc.authController.ChangeData)
	router.POST("/get-info", rc.authController.GetInfoByToken)
}
