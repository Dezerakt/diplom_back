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
		{
			Model:    gorm.Model{},
			SingerID: 2,
			Name:     "Men I Trust",
			Count:    13,
			Price:    222,
			ImageURL: "https://i.discogs.com/VC7Yb2cz67bW6Xy1Qz443jbwmRQvbSbF--h6Abzs18M/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTEwMTc2/NTU4LTE0OTg1MTM0/MDMtNDA2Ny5qcGVn.jpeg",
		},
		{
			Model:    gorm.Model{},
			SingerID: 3,
			Name:     "Welfare Jazz",
			Count:    14,
			Price:    123,
			ImageURL: "https://i.discogs.com/pWU-4LoNMgcf0IxlhJMroJU3K4hZPn9dz-t19X3sPyc/rs:fit/g:sm/q:90/h:593/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTE2ODA3/Nzg1LTE2MTAyMDM1/ODgtNTYyMi5qcGVn.jpeg",
		},
	})
}
