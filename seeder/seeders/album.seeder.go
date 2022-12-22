package seeders

import (
	"diplom_back/models"
	"gorm.io/gorm"
)

type AlbumSeeder struct{}

func (obj AlbumSeeder) Seed(DB *gorm.DB) {
	DB.Create([]models.Album{
		{
			Model:    gorm.Model{},
			SingerID: 1,
			Name:     "Mezmerize",
			Count:    12,
			Price:    123,
			ImageURL: "https://i.discogs.com/NZ1d7Rwr8Ut_mb7CU62fAUadxOUAP8H7xqOJ1nwH9Tw/rs:fit/g:sm/q:90/h:600/w:596/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTEyNjQy/NzAzLTE1NDAwNjQ1/NjItMTE2Ni5qcGVn.jpeg",
		},
	})
}
