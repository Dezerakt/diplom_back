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
					Model: gorm.Model{},
					Name:  "Soldier Side - Intro",
				},
				{
					Model: gorm.Model{},
					Name:  "B.Y.O.B.",
				},
				{
					Model: gorm.Model{},
					Name:  "Revenga",
				},
				{
					Model: gorm.Model{},
					Name:  "Cigaro",
				},
				{
					Model: gorm.Model{},
					Name:  "Radio/Video",
				},
				{
					Model: gorm.Model{},
					Name:  "This Cocaine Makes Me Feel Like I'm On This Song",
				},
				{
					Model: gorm.Model{},
					Name:  "Violent Pornography",
				},
				{
					Model: gorm.Model{},
					Name:  "Question!",
				},
				{
					Model: gorm.Model{},
					Name:  "Sad Statue",
				},
				{
					Model: gorm.Model{},
					Name:  "Old School Hollywood",
				},
				{
					Model: gorm.Model{},
					Name:  "Lost In Hollywood",
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
					Model: gorm.Model{},
					Name:  "Plain View",
				},
				{
					Model: gorm.Model{},
					Name:  "Morse Code",
				},
				{
					Model: gorm.Model{},
					Name:  "Humming Man",
				},
				{
					Model: gorm.Model{},
					Name:  "Stay True",
				},
				{
					Model: gorm.Model{},
					Name:  "Lauren",
				},
				{
					Model: gorm.Model{},
					Name:  "Again",
				},
				{
					Model: gorm.Model{},
					Name:  "Quiet",
				},
				{
					Model: gorm.Model{},
					Name:  "Break For Lovers",
				},
				{
					Model: gorm.Model{},
					Name:  "Aquarelle",
				},
				{
					Model: gorm.Model{},
					Name:  "Yes",
				},
				{
					Model: gorm.Model{},
					Name:  "Middle Points",
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
					Model: gorm.Model{},
					Name:  "Ain't Nice",
				},
				{
					Model: gorm.Model{},
					Name:  "Cold Play",
				},
				{
					Model: gorm.Model{},
					Name:  "Toad",
				},
				{
					Model: gorm.Model{},
					Name:  "This Old Dog",
				},
				{
					Model: gorm.Model{},
					Name:  "Into The Sun",
				},
				{
					Model: gorm.Model{},
					Name:  "Creatures",
				},
				{
					Model: gorm.Model{},
					Name:  "Best In Show II",
				},
				{
					Model: gorm.Model{},
					Name:  "Secret Canine Agent",
				},
				{
					Model: gorm.Model{},
					Name:  "I Feel Alive",
				},
				{
					Model: gorm.Model{},
					Name:  "Girls & Boys",
				},
				{
					Model: gorm.Model{},
					Name:  "To The Country",
				},
				{
					Model: gorm.Model{},
					Name:  "In Spite Of Ourselves",
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
					AlbumID: 3,
					URL:     "https://i.discogs.com/JZmwt6pfZyORIJ7Bharw38bJt_RQLewpJj5wHGrcWk0/rs:fit/g:sm/q:90/h:553/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTE2ODE4/MjQ2LTE2MTA1NDIw/MTctMzUzMi5qcGVn.jpeg",
				},
				{
					Model:   gorm.Model{},
					AlbumID: 3,
					URL:     "https://i.discogs.com/Rc-_Mi2Mbw3WVgDqAI7KqDVhebv8k9-evfaYn_wLWas/rs:fit/g:sm/q:90/h:450/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTE2ODE4/MjQ2LTE2MTM1OTMw/MjAtODY5OC5qcGVn.jpeg",
				},
			},
		},
		{
			Model:       gorm.Model{},
			SingerID:    4,
			SingerName:  "Kanye West",
			Name:        "Get Well Soon...",
			Count:       23,
			Price:       144,
			ImageURL:    "https://i.discogs.com/kcw9oytio_dauy3oR-GyWLbFRvKSSMZmP0rVVU8_WfQ/rs:fit/g:sm/q:90/h:600/w:597/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExOTQx/NDU1LTE1MjUxNTYz/NzEtNjIxOS5qcGVn.jpeg",
			ReleaseYear: 2016,
			Genre:       "Thug Rap",
			Songs: []models.Song{
				{
					Model: gorm.Model{},
					Name:  "Intro",
				},
				{
					Model: gorm.Model{},
					Name:  "Live From Irving Plaza, NY",
				},
				{
					Model: gorm.Model{},
					Name:  "Guess Who's Back? Freestyle",
				},
				{
					Model: gorm.Model{},
					Name:  "Jesus Walks",
				},
				{
					Model: gorm.Model{},
					Name:  "Through The Wire",
				},
				{
					Model: gorm.Model{},
					Name:  "2 Words",
				},
				{
					Model: gorm.Model{},
					Name:  "Show Go On",
				},
				{
					Model: gorm.Model{},
					Name:  "Champions",
				},
				{
					Model: gorm.Model{},
					Name:  "Live From Tweeter Center, IL",
				},
				{
					Model: gorm.Model{},
					Name:  "The Bounce",
				},
				{
					Model: gorm.Model{},
					Name:  "Poppin' Tags",
				},
				{
					Model: gorm.Model{},
					Name:  "A Million Freestyle",
				},
			},
			Images: []models.Image{
				{
					Model:   gorm.Model{},
					AlbumID: 4,
					URL:     "https://i.discogs.com/vLsg_cbhW5JXOxac_4A7kaxwXdllZP7nCvWfH330LrA/rs:fit/g:sm/q:90/h:600/w:598/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExOTQx/NDU1LTE1MjUxNTYz/NzQtMzMzMC5qcGVn.jpeg",
				},
				{
					Model:   gorm.Model{},
					AlbumID: 4,
					URL:     "https://i.discogs.com/NiipCdJCxeUpdr36N2eY8AK8BsVL66e544LuBtH56rI/rs:fit/g:sm/q:90/h:600/w:599/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTExOTQx/NDU1LTE1MjUxNTU4/NzUtNTE2My5qcGVn.jpeg",
				},
			},
		},
		{
			Model:       gorm.Model{},
			SingerID:    1,
			SingerName:  "System Of A Down",
			Name:        "Steal This Album!",
			Count:       13,
			Price:       155,
			ImageURL:    "https://i.discogs.com/HZLPl969-R2mhkpDKqw4VIemE7PGUwJdudLtnCQV2nw/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTc0Mjgz/NS0xNTM2ODk3MTA2/LTY2NjAucG5n.jpeg",
			ReleaseYear: 2002,
			Genre:       "Heavy Metal",
			Songs: []models.Song{
				{
					Model: gorm.Model{},
					Name:  "Chic 'N' Stu",
				},
				{
					Model: gorm.Model{},
					Name:  "Innervision",
				},
				{
					Model: gorm.Model{},
					Name:  "Bubbles",
				},
				{
					Model: gorm.Model{},
					Name:  "Boom!",
				},
				{
					Model: gorm.Model{},
					Name:  "Nüguns",
				},
				{
					Model: gorm.Model{},
					Name:  "A.D.D.",
				},
				{
					Model: gorm.Model{},
					Name:  "Mr. Jack",
				},
				{
					Model: gorm.Model{},
					Name:  "I-E-A-I-A-I-O",
				},
				{
					Model: gorm.Model{},
					Name:  "36",
				},
				{
					Model: gorm.Model{},
					Name:  "Pictures",
				},
				{
					Model: gorm.Model{},
					Name:  "Highway Song",
				},
			},
			Images: []models.Image{
				{
					Model:   gorm.Model{},
					AlbumID: 5,
					URL:     "https://i.discogs.com/3oYMz3pC2qAmU0bcULC7TJeVPfKN6100RY7PYTCwi14/rs:fit/g:sm/q:90/h:472/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTc0Mjgz/NS0xNTM2ODk3MTA5/LTQyMjMucG5n.jpeg",
				},
				{
					Model:   gorm.Model{},
					AlbumID: 5,
					URL:     "https://i.discogs.com/-Og4oJPa6oIxday755NBVtCidVWgzGLI8aMICoxV3SI/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTc0Mjgz/NS0xNTM2ODk3MTEx/LTE2MDMucG5n.jpeg",
				},
				{
					Model:   gorm.Model{},
					AlbumID: 5,
					URL:     "https://i.discogs.com/xvbmPPqPUdan9HZZdEGMbGe5PkifdN-Xqnvuq8k4ODA/rs:fit/g:sm/q:90/h:472/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTc0Mjgz/NS0xNTM2ODk3MTEx/LTQ1MjMucG5n.jpeg",
				},
			},
		},
		{
			Model:       gorm.Model{},
			SingerID:    1,
			SingerName:  "System Of A Down",
			Name:        "Toxicity",
			Count:       19,
			Price:       143,
			ImageURL:    "https://i.discogs.com/ZNHCc_2QjdbFftseh84ploWB6C0dx2YMnkVmcszXCYA/rs:fit/g:sm/q:90/h:592/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTYwODM0/Ni0xMTkxMDMyMjU0/LmpwZWc.jpeg",
			ReleaseYear: 2003,
			Genre:       "Alternative Rock",
			Songs: []models.Song{
				{
					Model: gorm.Model{},
					Name:  "Prison Song",
				},
				{
					Model: gorm.Model{},
					Name:  "Needles",
				},
				{
					Model: gorm.Model{},
					Name:  "Deer Dance",
				},
				{
					Model: gorm.Model{},
					Name:  "Jet Pilot",
				},
				{
					Model: gorm.Model{},
					Name:  "X",
				},
				{
					Model: gorm.Model{},
					Name:  "Chop Suey!",
				},
				{
					Model: gorm.Model{},
					Name:  "Bounce",
				},
				{
					Model: gorm.Model{},
					Name:  "Forest",
				},
				{
					Model: gorm.Model{},
					Name:  "ATWA",
				},
				{
					Model: gorm.Model{},
					Name:  "Science",
				},
				{
					Model: gorm.Model{},
					Name:  "Shimmy",
				},
				{
					Model: gorm.Model{},
					Name:  "Toxicity",
				},
				{
					Model: gorm.Model{},
					Name:  "Psycho",
				},
				{
					Model: gorm.Model{},
					Name:  "Aerials",
				},
			},
			Images: []models.Image{
				{
					Model:   gorm.Model{},
					AlbumID: 6,
					URL:     "https://i.discogs.com/FOHcKyCmvrQfiUFlxJDnRAFR9vdkt1MkMWD1moR11Vs/rs:fit/g:sm/q:90/h:595/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTYwODM0/Ni0xMTkxMDMyMzI0/LmpwZWc.jpeg",
				},
				{
					Model:   gorm.Model{},
					AlbumID: 6,
					URL:     "https://i.discogs.com/XEy4HzZO7Cb4KyHMyHTjA3YuxHFJYBatXoiyhL-rZ10/rs:fit/g:sm/q:90/h:549/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTYwODM0/Ni0xNDM3MzE4NDg3/LTIyNjkuanBlZw.jpeg",
				},
				{
					Model:   gorm.Model{},
					AlbumID: 6,
					URL:     "https://i.discogs.com/w5R2Ho44ZZcCT8PkkjuhGGZkOB7eXhtum3JYwfMyPhc/rs:fit/g:sm/q:90/h:600/w:586/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTYwODM0/Ni0xNDM3MzE4Mzk3/LTg4MTQuanBlZw.jpeg",
				},
			},
		},
		{
			Model:       gorm.Model{},
			SingerID:    4,
			SingerName:  "Kanye West",
			Name:        "Late Registration",
			Count:       42,
			Price:       155,
			ImageURL:    "https://i.discogs.com/VaGAJ7bFhiP1JY5nrQDFGGw2mmbU7OH0xm_j9msSaNM/rs:fit/g:sm/q:90/h:592/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTQ5NDE2/OS0xNTAxMjY5NDM1/LTMxNzIuanBlZw.jpeg",
			ReleaseYear: 2005,
			Genre:       "Alternative Rock",
			Songs: []models.Song{
				{
					Model: gorm.Model{},
					Name:  "Wake Up Mr. West",
				},
				{
					Model: gorm.Model{},
					Name:  "Heard 'Em Say",
				},
				{
					Model: gorm.Model{},
					Name:  "Touch The Sky",
				},
				{
					Model: gorm.Model{},
					Name:  "Gold Digger",
				},
				{
					Model: gorm.Model{},
					Name:  "Skit #1",
				},
				{
					Model: gorm.Model{},
					Name:  "Drive Slow",
				},
				{
					Model: gorm.Model{},
					Name:  "My Way Home",
				},
				{
					Model: gorm.Model{},
					Name:  "Crack Music",
				},
				{
					Model: gorm.Model{},
					Name:  "Roses",
				},
				{
					Model: gorm.Model{},
					Name:  "Bring Me Down",
				},
				{
					Model: gorm.Model{},
					Name:  "Addiction",
				},
				{
					Model: gorm.Model{},
					Name:  "Skit #2",
				},
				{
					Model: gorm.Model{},
					Name:  "Diamonds From Sierra Leone",
				},
				{
					Model: gorm.Model{},
					Name:  "We Major",
				},
			},
			Images: []models.Image{
				{
					Model:   gorm.Model{},
					AlbumID: 7,
					URL:     "https://i.discogs.com/fhppOFuMlLIW8IuA37xoaLwKGvV4Z4rneITrmzVdHb8/rs:fit/g:sm/q:90/h:600/w:589/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTQ5NDE2/OS0xNTExMzc3OTc3/LTg4MTguanBlZw.jpeg",
				},
				{
					Model:   gorm.Model{},
					AlbumID: 7,
					URL:     "https://i.discogs.com/Ud63CnLeN5VGYhhmPEwusxuQAt2iK4UGw1wf1JMW_kQ/rs:fit/g:sm/q:90/h:551/w:541/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTQ5NDE2/OS0xNTExMzc3OTc2/LTc2NjAuanBlZw.jpeg",
				},
				{
					Model:   gorm.Model{},
					AlbumID: 7,
					URL:     "https://i.discogs.com/ksZxBa8Qqpc1e1Njp_lVshxIRsBFFIiDWnGTrU9n1bA/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTQ5NDE2/OS0xNTAxMjY5NDQ4/LTg5MDcuanBlZw.jpeg",
				},
			},
		},
		{
			Model:       gorm.Model{},
			SingerID:    3,
			SingerName:  "Viagra Boys",
			Name:        "Street Worms",
			Count:       36,
			Price:       142,
			ImageURL:    "https://i.discogs.com/IxzqQe47S6NlvqLpeuFs4DwNLj10v2_dWzTa5BMBOwQ/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTEyNDkw/ODUyLTE1MzYzMzIw/NTUtOTk2NS5qcGVn.jpeg",
			ReleaseYear: 2018,
			Genre:       "Indie Rock",
			Songs: []models.Song{
				{
					Model: gorm.Model{},
					Name:  "Down In The Basement",
				},
				{
					Model: gorm.Model{},
					Name:  "Slow Learner",
				},
				{
					Model: gorm.Model{},
					Name:  "Sports",
				},
				{
					Model: gorm.Model{},
					Name:  "Best In Show",
				},
				{
					Model: gorm.Model{},
					Name:  "Just Like You",
				},
				{
					Model: gorm.Model{},
					Name:  "Shrimp Shack",
				},
				{
					Model: gorm.Model{},
					Name:  "Frogstrap",
				},
				{
					Model: gorm.Model{},
					Name:  "Worms",
				},
				{
					Model: gorm.Model{},
					Name:  "Amphetanarchy",
				},
			},
			Images: []models.Image{
				{
					Model:   gorm.Model{},
					AlbumID: 8,
					URL:     "https://i.discogs.com/aoej-T3oHwO0Ew_1TbR_-LW-XAnFjdpH8eYWrvmKfmk/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTEyNDkw/ODUyLTE1Mzk5NDIw/NjMtNjAwNC5qcGVn.jpeg",
				},
				{
					Model:   gorm.Model{},
					AlbumID: 8,
					URL:     "https://i.discogs.com/4Qoj2o7isfddvK4wP9CUO74ix3yGRCls2G7CypAc2cw/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTEyNDkw/ODUyLTE1Mzk5NDIw/NjAtNjQ4OS5qcGVn.jpeg",
				},
				{
					Model:   gorm.Model{},
					AlbumID: 8,
					URL:     "https://i.discogs.com/HuOayOiQ4bOyPufQoJH4vM59SqO8gAp-Pu4-vqhmtNc/rs:fit/g:sm/q:90/h:596/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTEyNDkw/ODUyLTE1Mzk5NDIw/NTgtMTUzNC5qcGVn.jpeg",
				},
			},
		},
		{
			Model:       gorm.Model{},
			SingerID:    2,
			SingerName:  "Men I Trust",
			Name:        "Untourable Album",
			Count:       22,
			Price:       136,
			ImageURL:    "https://i.discogs.com/yD98cyW4w3wi3ltuaEH23FBdZS8ewwvdcqB5vPDl1FQ/rs:fit/g:sm/q:90/h:598/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTIwOTAw/Mzc3LTE2NDgxOTk2/NjYtNzU2Mi5qcGVn.jpeg",
			ReleaseYear: 2021,
			Genre:       "Indie Pop",
			Songs: []models.Song{
				{
					Model: gorm.Model{},
					Name:  "Organon",
				},
				{
					Model: gorm.Model{},
					Name:  "Oh Dove",
				},
				{
					Model: gorm.Model{},
					Name:  "Sugar",
				},
				{
					Model: gorm.Model{},
					Name:  "Sorbitol",
				},
				{
					Model: gorm.Model{},
					Name:  "Tree Among Shrubs",
				},
				{
					Model: gorm.Model{},
					Name:  "Before Dawn",
				},
				{
					Model: gorm.Model{},
					Name:  "Serenade Of Water",
				},
				{
					Model: gorm.Model{},
					Name:  "5AM Waltz",
				},
				{
					Model: gorm.Model{},
					Name:  "Always Lone",
				},
			},
			Images: []models.Image{
				{
					Model:   gorm.Model{},
					AlbumID: 9,
					URL:     "https://i.discogs.com/YYBuvhLIf_s-egPqHkG1EbTor9jXfIQK3rGNWWhCrWM/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTIwOTAw/Mzc3LTE2NTY0NTc4/MDQtNDE4OC5qcGVn.jpeg",
				},
				{
					Model:   gorm.Model{},
					AlbumID: 9,
					URL:     "https://i.discogs.com/x7iHGR7Wi6WCIDMp3GzlPkR7cmLL9GS5KuVyBrmSsaM/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTIwOTAw/Mzc3LTE2NTY0NTc4/MDQtNDU5Ny5qcGVn.jpeg",
				},
				{
					Model:   gorm.Model{},
					AlbumID: 9,
					URL:     "https://i.discogs.com/q0ySJ-nSiA-pU_jOgyY2Kw30XMr4tSbj-cHUjzdmtSc/rs:fit/g:sm/q:90/h:596/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTIwOTAw/Mzc3LTE2NDgyMDAx/MTItNTQ0OS5qcGVn.jpeg",
				},
			},
		},
		{
			Model:       gorm.Model{},
			SingerID:    5,
			SingerName:  "Limp Bizkit",
			Name:        "Gold Cobra",
			Count:       32,
			Price:       144,
			ImageURL:    "https://i.discogs.com/BmLIpVpu7-TB0daVvRKVnJpJKNVVndHyIEmEFQtzbl4/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTQzOTAz/MjQtMTU2MjEzMDk2/Mi04OTI5LmpwZWc.jpeg",
			ReleaseYear: 2021,
			Genre:       "Nu Metal",
			Songs: []models.Song{
				{
					Model: gorm.Model{},
					Name:  "Introbra",
				},
				{
					Model: gorm.Model{},
					Name:  "Bring It Back",
				},
				{
					Model: gorm.Model{},
					Name:  "Gold Cobra",
				},
				{
					Model: gorm.Model{},
					Name:  "Shark Attack",
				},
				{
					Model: gorm.Model{},
					Name:  "Get A Life / Interlude 1",
				},
				{
					Model: gorm.Model{},
					Name:  "Shotgun",
				},
				{
					Model: gorm.Model{},
					Name:  "Douche Bag",
				},
				{
					Model: gorm.Model{},
					Name:  "Walking Away",
				},
				{
					Model: gorm.Model{},
					Name:  "Loser / Interlude 2",
				},
			},
			Images: []models.Image{
				{
					Model:   gorm.Model{},
					AlbumID: 10,
					URL:     "https://i.discogs.com/mkkX-8O4Ew_UOTQDy6WzDNj2vKJwzRe1aBkZD4AuKH0/rs:fit/g:sm/q:90/h:461/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTQzOTAz/MjQtMTU3MDA0NTQz/My0xNzYyLmpwZWc.jpeg",
				},
				{
					Model:   gorm.Model{},
					AlbumID: 10,
					URL:     "https://i.discogs.com/4hNGFfElzL-sVRN7e1xDWl8R4OhsvWHO8TR6QAVHJQk/rs:fit/g:sm/q:90/h:599/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTQzOTAz/MjQtMTU3MDA0NTQz/My00MTUxLmpwZWc.jpeg",
				},
				{
					Model:   gorm.Model{},
					AlbumID: 10,
					URL:     "https://i.discogs.com/6L5JX2MJI56DnDoyX0gMXKpM8cUWb5YSy7_L1qCGDVc/rs:fit/g:sm/q:90/h:600/w:599/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTQzOTAz/MjQtMTY0NDQxNTc3/Mi0xMDU4LmpwZWc.jpeg",
				},
			},
		},
		{
			Model:       gorm.Model{},
			SingerID:    6,
			SingerName:  "Nirvana",
			Name:        "Nevermind",
			Count:       23,
			Price:       233,
			ImageURL:    "https://i.discogs.com/XoAXkLe03Vk0Kt-oQqw7V9uW5nBzQdNvZd0zar5mOiQ/rs:fit/g:sm/q:90/h:592/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTM2NzA4/NC0xMjYzMDk1NTUz/LmpwZWc.jpeg",
			ReleaseYear: 1991,
			Genre:       "Grunge",
			Songs: []models.Song{
				{
					Model: gorm.Model{},
					Name:  "Smells Like Teen Spirit",
				},
				{
					Model: gorm.Model{},
					Name:  "In Bloom",
				},
				{
					Model: gorm.Model{},
					Name:  "Come As You Are",
				},
				{
					Model: gorm.Model{},
					Name:  "Breed",
				},
				{
					Model: gorm.Model{},
					Name:  "Lithium",
				},
				{
					Model: gorm.Model{},
					Name:  "Polly",
				},
				{
					Model: gorm.Model{},
					Name:  "Territorial Pissings",
				},
				{
					Model: gorm.Model{},
					Name:  "Drain You",
				},
				{
					Model: gorm.Model{},
					Name:  "Lounge Act",
				},
			},
			Images: []models.Image{
				{
					Model:   gorm.Model{},
					AlbumID: 11,
					URL:     "https://i.discogs.com/n2nXzjI6p5xuv_Dn_wlhLmCa2R7n2TbrpjM8AowSxOQ/rs:fit/g:sm/q:90/h:466/w:581/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTM2NzA4/NC0xNDk2OTA2ODU1/LTYzMTguanBlZw.jpeg",
				},
				{
					Model:   gorm.Model{},
					AlbumID: 11,
					URL:     "https://i.discogs.com/fUJDG7yC36gWOARmYRT4P-MdoauzogmlgO75t7YTDlo/rs:fit/g:sm/q:90/h:597/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTM2NzA4/NC0xNDE2NjAwMTcx/LTc2OTEuanBlZw.jpeg",
				},
				{
					Model:   gorm.Model{},
					AlbumID: 11,
					URL:     "https://i.discogs.com/4ZxF0pNtR8QYaAkxM81cLexExaAedD9xrvknkeqGOq4/rs:fit/g:sm/q:90/h:600/w:475/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTM2NzA4/NC0xNDkwNDQ3Mzgz/LTY0MjIuanBlZw.jpeg",
				},
			},
		},
		{
			Model:       gorm.Model{},
			SingerID:    5,
			SingerName:  "Limp Bizkit",
			Name:        "Three Dollar Bill, Yall$",
			Count:       23,
			Price:       233,
			ImageURL:    "https://i.discogs.com/Obk6L165Pq-99y10u0oNbp4dfA3G55rCFTiENWYV7bU/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTQwNjA2/MC0xMjgyMTY3MTUw/LmpwZWc.jpeg",
			ReleaseYear: 1997,
			Genre:       "Hard Rock",
			Songs: []models.Song{
				{
					Model: gorm.Model{},
					Name:  "Intro",
				},
				{
					Model: gorm.Model{},
					Name:  "Pollution",
				},
				{
					Model: gorm.Model{},
					Name:  "Counterfeit",
				},
				{
					Model: gorm.Model{},
					Name:  "Stuck",
				},
				{
					Model: gorm.Model{},
					Name:  "Nobody ♥'s Me",
				},
				{
					Model: gorm.Model{},
					Name:  "Sour",
				},
				{
					Model: gorm.Model{},
					Name:  "Stalemate",
				},
				{
					Model: gorm.Model{},
					Name:  "Clunk",
				},
				{
					Model: gorm.Model{},
					Name:  "Faith",
				},
			},
			Images: []models.Image{
				{
					Model:   gorm.Model{},
					AlbumID: 12,
					URL:     "https://i.discogs.com/0bupyLmcUdMcj_lm4yzLAJ-a6jtE4q_vf3UZOKTMlcM/rs:fit/g:sm/q:90/h:589/w:592/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTQwNjA2/MC0xMTY1MjA0NDIy/LmpwZWc.jpeg	",
				},
				{
					Model:   gorm.Model{},
					AlbumID: 12,
					URL:     "https://i.discogs.com/LAk4nIDrISekbtJmDAYOOE6tWp3a10PLtMWLqL912L0/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTQwNjA2/MC0xMzg4MjIwMjM3/LTMwNjAuanBlZw.jpeg",
				},
				{
					Model:   gorm.Model{},
					AlbumID: 12,
					URL:     "https://i.discogs.com/FUsXNV7C8TSqEgnwAjoTLf6x-GNYjSTuLdpDcNUARoA/rs:fit/g:sm/q:90/h:471/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTQwNjA2/MC0xMjgyMTY3MjYz/LmpwZWc.jpeg",
				},
			},
		},
		{
			Model:       gorm.Model{},
			SingerID:    6,
			SingerName:  "Nirvana",
			Name:        "Bleach",
			Count:       12,
			Price:       534,
			ImageURL:    "https://i.discogs.com/WTTg74ExM5Aq0hkwV_6CWO6LCOh-2mepiT6UOVUh238/rs:fit/g:sm/q:90/h:592/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTI3OTU1/NTQtMTMwMTM3OTg3/Mi5qcGVn.jpeg",
			ReleaseYear: 1989,
			Genre:       "Grunge",
			Songs: []models.Song{
				{
					Model: gorm.Model{},
					Name:  "Blew",
				},
				{
					Model: gorm.Model{},
					Name:  "Floyd The Barber",
				},
				{
					Model: gorm.Model{},
					Name:  "About A Girl",
				},
				{
					Model: gorm.Model{},
					Name:  "School",
				},
				{
					Model: gorm.Model{},
					Name:  "Love Buzz",
				},
				{
					Model: gorm.Model{},
					Name:  "Paper Cuts",
				},
				{
					Model: gorm.Model{},
					Name:  "Negative Creep",
				},
				{
					Model: gorm.Model{},
					Name:  "Scoff",
				},
				{
					Model: gorm.Model{},
					Name:  "Swap Meet",
				},
				{
					Model: gorm.Model{},
					Name:  "Mr. Moustache",
				},
				{
					Model: gorm.Model{},
					Name:  "Sifting",
				},
			},
			Images: []models.Image{
				{
					Model:   gorm.Model{},
					AlbumID: 12,
					URL:     "https://i.discogs.com/0bupyLmcUdMcj_lm4yzLAJ-a6jtE4q_vf3UZOKTMlcM/rs:fit/g:sm/q:90/h:589/w:592/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTQwNjA2/MC0xMTY1MjA0NDIy/LmpwZWc.jpeg	",
				},
				{
					Model:   gorm.Model{},
					AlbumID: 12,
					URL:     "https://i.discogs.com/LAk4nIDrISekbtJmDAYOOE6tWp3a10PLtMWLqL912L0/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTQwNjA2/MC0xMzg4MjIwMjM3/LTMwNjAuanBlZw.jpeg",
				},
				{
					Model:   gorm.Model{},
					AlbumID: 12,
					URL:     "https://i.discogs.com/FUsXNV7C8TSqEgnwAjoTLf6x-GNYjSTuLdpDcNUARoA/rs:fit/g:sm/q:90/h:471/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9SLTQwNjA2/MC0xMjgyMTY3MjYz/LmpwZWc.jpeg",
				},
			},
		},
	})
}
