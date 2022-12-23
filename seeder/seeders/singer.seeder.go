package seeders

import (
	"diplom_back/models"
	"gorm.io/gorm"
	"time"
)

type SingerSeeder struct {
}

func (obj SingerSeeder) Seed(DB *gorm.DB) {
	DB.Create([]models.Singer{
		{
			Model:     gorm.Model{ID: 1},
			Name:      "System Of A Down",
			BirthDate: time.Date(2003, 12, 03, 0, 0, 0, 0, time.Local),
		},
		{
			Model:     gorm.Model{ID: 2},
			BirthDate: time.Date(2003, 03, 04, 0, 0, 0, 0, time.Local),
			Name:      "Men I Trust",
		},
		{
			Model:     gorm.Model{ID: 3},
			BirthDate: time.Date(2003, 03, 04, 0, 0, 0, 0, time.Local),
			Name:      "Viagra Boys",
		},
	})
}
