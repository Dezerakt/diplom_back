package main

import (
	"diplom_back/controllers"
	"diplom_back/initializers"
	"diplom_back/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

var (
	server              *gin.Engine
	AuthController      controllers.AuthController
	AuthRouteController routes.AuthRouteController

	AlbumController      controllers.AlbumController
	AlbumRouteController routes.AlbumRouteController
	/*
		UserController      controllers.UserController
		UserRouteController routes.UserRouteController

		PostController      controllers.PostController
		PostRouteController routes.PostRouteController*/
)

func init() {
	config, err := initializers.LoadConfig(".")

	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	AuthController = controllers.NewAuthController(initializers.DB)
	AuthRouteController = routes.NewAuthRouteController(AuthController)
	AlbumController = controllers.NewAlbumController(initializers.DB)
	AlbumRouteController = routes.NewAlbumRouteController(AlbumController)

	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8080", config.ClientOrigin}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")

	AuthRouteController.AuthRoute(router)
	AlbumRouteController.AlbumRoute(router)
	log.Fatal(server.Run(":" + config.ServerPort))

}
