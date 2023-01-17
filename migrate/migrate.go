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
	err := initializers.DB.Migrator().DropTable(
		models.Song{},
		models.Image{},
		models.Singer{},
		models.Album{},
		models.Token{},
		models.User{},
		models.Order{},
		models.Cart{})
	if err != nil {
		log.Fatal("Error while removing data from database", err)
	}
	err = initializers.DB.AutoMigrate(
		models.Song{},
		models.Image{},
		models.Singer{},
		models.Album{},
		models.Token{},
		models.User{},
		models.Order{},
		models.Cart{})
	if err != nil {
		log.Fatal("Could not complete migrations", err)
	}
	fmt.Println("Migrations complete")
}
