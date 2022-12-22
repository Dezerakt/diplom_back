package main

import (
	"diplom_back/initializers"
	"diplom_back/seeder/seeders"
	"fmt"
	"log"
)

var (
	AlbumSeeder  seeders.AlbumSeeder
	SingerSeeder seeders.SingerSeeder
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variable")
	}

	initializers.ConnectDB(&config)
}

func main() {
	SingerSeeder.Seed(initializers.DB)
	AlbumSeeder.Seed(initializers.DB)

	fmt.Println("Seeding completed")
}
