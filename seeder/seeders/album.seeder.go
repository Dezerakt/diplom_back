package seeders

import (
	"diplom_back/models"
	"gorm.io/gorm"
)

type AlbumSeeder struct{}

func (obj AlbumSeeder) Seed(DB *gorm.DB) {
	DB.Create([]models.Album{
		{
			Model:       gorm.Model{},
			SingerID:    1,
			SingerName:  "System Of A Down",
			Name:        "Mezmerize",
			Count:       12,
			Price:       123,
			ImageURL:    "https://i.discogs.com/NZ1d7Rwr8Ut_mb7CU62fAUadxOUAP8H7xqOJ1nwH9Tw/rs:fit/g:sm/q:90/h:600/w:596/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTEyNjQy/NzAzLTE1NDAwNjQ1/NjItMTE2Ni5qcGVn.jpeg",
			ReleaseYear: 2005,
			Genre:       "Rock",
			Songs: []models.Song{
				{
					Model:         gorm.Model{},
					Duration:      "1:03",
					Name:          "Soldier Side - Intro",
					SpotifyWidget: "",
				},
				{
					Model:         gorm.Model{},
					Duration:      "4:15",
					Name:          "B.Y.O.B.",
					SpotifyWidget: "",
				},
				{
					Model:         gorm.Model{},
					Duration:      "3:48",
					Name:          "Revenga",
					SpotifyWidget: "",
				},
				{
					Model:         gorm.Model{},
					Duration:      "2:11",
					Name:          "Cigaro",
					SpotifyWidget: "",
				},
				{
					Model:         gorm.Model{},
					Duration:      "4:09",
					Name:          "Radio/Video",
					SpotifyWidget: "",
				},
				{
					Model:         gorm.Model{},
					Duration:      "2:08",
					Name:          "This Cocaine Makes Me Feel Like I'm On This Song",
					SpotifyWidget: "",
				},
				{
					Model:         gorm.Model{},
					Duration:      "3:31",
					Name:          "Violent Pornography",
					SpotifyWidget: "",
				},
				{
					Model:         gorm.Model{},
					Duration:      "3:20",
					Name:          "Question!",
					SpotifyWidget: "",
				},
				{
					Model:         gorm.Model{},
					Duration:      "3:25",
					Name:          "Sad Statue",
					SpotifyWidget: "",
				},
				{
					Model:         gorm.Model{},
					Duration:      "2:56",
					Name:          "Old School Hollywood",
					SpotifyWidget: "",
				},
				{
					Model:         gorm.Model{},
					Duration:      "5:20",
					Name:          "Lost In Hollywood",
					SpotifyWidget: "",
				},
			},
			Images: []models.Image{
				{
					Model:   gorm.Model{},
					AlbumID: 1,
					URL:     "https://i.discogs.com/kw9LTBdEe5A_iTuXrIX24AJq3kwha2c1dWTuaT5O4mo/rs:fit/g:sm/q:90/h:294/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTQ3OTQz/MC0xNTIzNTMxNzU5/LTk5OTguanBlZw.jpeg",
				},
				{
					Model:   gorm.Model{},
					AlbumID: 1,
					URL:     "https://i.discogs.com/2elnj_rxqn_B_VAKrXKhzzpS1tDskHkKLp9hnEOtS5s/rs:fit/g:sm/q:90/h:590/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTQ3OTQz/MC0xNDI5MjIzOTA0/LTQ4MDAuanBlZw.jpeg",
				},
				{
					Model:   gorm.Model{},
					AlbumID: 1,
					URL:     "https://i.discogs.com/K6srTN8RYMQ5uUKBQS3p20NdDZRdyppUOmf8n-uwhak/rs:fit/g:sm/q:90/h:592/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTQ3OTQz/MC0xNDI5MjIzOTA0/LTMwMTAuanBlZw.jpeg",
				},
			},
		},
		{
			Model:       gorm.Model{},
			SingerID:    2,
			SingerName:  "Men I Trust",
			Name:        "Men I Trust",
			Count:       13,
			Price:       222,
			ImageURL:    "https://i.discogs.com/VC7Yb2cz67bW6Xy1Qz443jbwmRQvbSbF--h6Abzs18M/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTEwMTc2/NTU4LTE0OTg1MTM0/MDMtNDA2Ny5qcGVn.jpeg",
			ReleaseYear: 2017,
			Genre:       "Rock",
			Songs: []models.Song{
				{
					Model:         gorm.Model{},
					Duration:      "1:03",
					Name:          "Soldier Side - Intro",
					SpotifyWidget: "",
				},
				{
					Model:         gorm.Model{},
					Duration:      "4:15",
					Name:          "B.Y.O.B.",
					SpotifyWidget: "",
				},
				{
					Model:         gorm.Model{},
					Duration:      "3:48",
					Name:          "Revenga",
					SpotifyWidget: "",
				},
				{
					Model:         gorm.Model{},
					Duration:      "2:11",
					Name:          "Cigaro",
					SpotifyWidget: "",
				},
				{
					Model:         gorm.Model{},
					Duration:      "4:09",
					Name:          "Radio/Video",
					SpotifyWidget: "",
				},
				{
					Model:         gorm.Model{},
					Duration:      "2:08",
					Name:          "This Cocaine Makes Me Feel Like I'm On This Song",
					SpotifyWidget: "",
				},
				{
					Model:         gorm.Model{},
					Duration:      "3:31",
					Name:          "Violent Pornography",
					SpotifyWidget: "",
				},
				{
					Model:         gorm.Model{},
					Duration:      "3:20",
					Name:          "Question!",
					SpotifyWidget: "",
				},
				{
					Model:         gorm.Model{},
					Duration:      "3:25",
					Name:          "Sad Statue",
					SpotifyWidget: "",
				},
				{
					Model:         gorm.Model{},
					Duration:      "2:56",
					Name:          "Old School Hollywood",
					SpotifyWidget: "",
				},
				{
					Model:         gorm.Model{},
					Duration:      "5:20",
					Name:          "Lost In Hollywood",
					SpotifyWidget: "",
				},
			},
			Images: []models.Image{
				{
					Model:   gorm.Model{},
					AlbumID: 2,
					URL:     "https://i.discogs.com/issTpFM1u6MvfolnvNhA0xbA_94f64OoozafINiPblI/rs:fit/g:sm/q:90/h:594/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTEwMTc2/NTU4LTE0OTI5MDQ0/OTctMjQzNS5qcGVn.jpeg",
				},
				{
					Model:   gorm.Model{},
					AlbumID: 2,
					URL:     "https://i.discogs.com/7MTEjQ-tL9IUsVP13fV2x8onryTSacI9OW4V-B5fv2I/rs:fit/g:sm/q:90/h:542/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTEwMTc2/NTU4LTE0OTI5MDQ0/OTctMjA1MS5qcGVn.jpeg",
				},
				{
					Model:   gorm.Model{},
					AlbumID: 2,
					URL:     "https://i.discogs.com/9sLx5F4KHCVzzmJet6YGxCnqsSv2g7QEqSd5N-mwbbU/rs:fit/g:sm/q:90/h:599/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTEwMTc2/NTU4LTE0OTI5MDQ0/OTctODI5MC5qcGVn.jpeg",
				},
			},
		},
		{
			Model:       gorm.Model{},
			SingerID:    3,
			SingerName:  "Viagra Boys",
			Name:        "Welfare Jazz",
			Count:       14,
			Price:       123,
			ImageURL:    "https://i.discogs.com/pWU-4LoNMgcf0IxlhJMroJU3K4hZPn9dz-t19X3sPyc/rs:fit/g:sm/q:90/h:593/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTE2ODA3/Nzg1LTE2MTAyMDM1/ODgtNTYyMi5qcGVn.jpeg",
			ReleaseYear: 2017,
			Genre:       "Rock",
			Songs: []models.Song{
				{
					Model:         gorm.Model{},
					Duration:      "1:03",
					Name:          "Soldier Side - Intro",
					SpotifyWidget: "",
				},
				{
					Model:         gorm.Model{},
					Duration:      "4:15",
					Name:          "B.Y.O.B.",
					SpotifyWidget: "",
				},
				{
					Model:         gorm.Model{},
					Duration:      "3:48",
					Name:          "Revenga",
					SpotifyWidget: "",
				},
				{
					Model:         gorm.Model{},
					Duration:      "2:11",
					Name:          "Cigaro",
					SpotifyWidget: "",
				},
				{
					Model:         gorm.Model{},
					Duration:      "4:09",
					Name:          "Radio/Video",
					SpotifyWidget: "",
				},
				{
					Model:         gorm.Model{},
					Duration:      "2:08",
					Name:          "This Cocaine Makes Me Feel Like I'm On This Song",
					SpotifyWidget: "",
				},
				{
					Model:         gorm.Model{},
					Duration:      "3:31",
					Name:          "Violent Pornography",
					SpotifyWidget: "",
				},
				{
					Model:         gorm.Model{},
					Duration:      "3:20",
					Name:          "Question!",
					SpotifyWidget: "",
				},
				{
					Model:         gorm.Model{},
					Duration:      "3:25",
					Name:          "Sad Statue",
					SpotifyWidget: "",
				},
				{
					Model:         gorm.Model{},
					Duration:      "2:56",
					Name:          "Old School Hollywood",
					SpotifyWidget: "",
				},
				{
					Model:         gorm.Model{},
					Duration:      "5:20",
					Name:          "Lost In Hollywood",
					SpotifyWidget: "",
				},
			},
			Images: []models.Image{
				{
					Model:   gorm.Model{},
					AlbumID: 3,
					URL:     "https://i.discogs.com/Abj590C6Waz7QEuHQmLZJrFEoXi_Sc5T-Dcp7jdiDyM/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTE2ODE4/MjQ2LTE2MTg3NDA2/NDktNzIwNy5qcGVn.jpeg",
				},
				{
					Model:   gorm.Model{},
					AlbumID: 1,
					URL:     "https://i.discogs.com/JZmwt6pfZyORIJ7Bharw38bJt_RQLewpJj5wHGrcWk0/rs:fit/g:sm/q:90/h:553/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTE2ODE4/MjQ2LTE2MTA1NDIw/MTctMzUzMi5qcGVn.jpeg",
				},
				{
					Model:   gorm.Model{},
					AlbumID: 1,
					URL:     "https://i.discogs.com/Rc-_Mi2Mbw3WVgDqAI7KqDVhebv8k9-evfaYn_wLWas/rs:fit/g:sm/q:90/h:450/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTE2ODE4/MjQ2LTE2MTM1OTMw/MjAtODY5OC5qcGVn.jpeg",
				},
			},
		},
	})
}
