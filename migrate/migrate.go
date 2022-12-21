package main

import (
	"diplom_back/initializers"
	"diplom_back/models"
	"fmt"
	"log"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variable")
	}

	initializers.ConnectDB(&config)
}

func main() {
	err := initializers.DB.AutoMigrate(models.Singer{}, models.Album{}, models.User{}, models.Order{})
	if err != nil {
		log.Fatal("Could not complete migrations")
	}
	fmt.Println("Migrations complete")
}
