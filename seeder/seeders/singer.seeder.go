package seeders

import (
	"diplom_back/models"
	"gorm.io/gorm"
)

type SingerSeeder struct {
}

func (obj SingerSeeder) Seed(DB *gorm.DB) {
	DB.Create([]models.Singer{
		{
			Model: gorm.Model{
				ID: 1,
			},
			Name: "System Of A Down",
		},
	})
}
