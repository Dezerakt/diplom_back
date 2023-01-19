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
			Model:       gorm.Model{ID: 1},
			Name:        "System Of A Down",
			BirthDate:   time.Date(2003, 12, 03, 0, 0, 0, 0, time.Local),
			ImageURL:    "https://i.discogs.com/zVFzCV5QLl45qSeG_whL2GHcJqAq6lmTHKitCU7iVt0/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9BLTg1ODg5/LTE2NDY2ODQ4Nzgt/Nzk0OS5qcGVn.jpeg",
			Description: "System of a Down, sometimes shortened to System and abbreviated as SOAD, is an Armenian-American heavy metal band from Glendale, California, formed in 1994. The band currently consists of Serj Tankian (lead vocals, keyboards), Daron Malakian (vocals, guitar), Shavo Odadjian (bass, backing vocals) and John Dolmayan (drums). ",
		},
		{
			Model:       gorm.Model{ID: 2},
			BirthDate:   time.Date(2003, 03, 04, 0, 0, 0, 0, time.Local),
			Name:        "Men I Trust",
			ImageURL:    "https://i.discogs.com/XJqvPDAdQ-cKlSAbmzM-wRzQ4wAAZN0hKKXzQxRg5jQ/rs:fit/g:sm/q:90/h:600/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9BLTM4Njk2/NDAtMTU4MDMzNzEx/Ni03Mzg4LmpwZWc.jpeg",
			Description: "Men I Trust is an indie pop band from Quebec City (Canada).\n\nFounding members Jessy and Dragos knew each other since their fourth year of high school. Back then (the mid-2000’s), Dragos had somewhat of a background in classical piano and Jessy mastered the dark and esoteric art of the metal guitar soloist.",
		},
		{
			Model:       gorm.Model{ID: 3},
			BirthDate:   time.Date(2003, 03, 04, 0, 0, 0, 0, time.Local),
			ImageURL:    "https://i.discogs.com/224J88x346xYKMgbiAXj1uvk7DXJFBAeqydv5iE8nY0/rs:fit/g:sm/q:90/h:500/w:500/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9BLTQ4NjEy/ODUtMTQ3MTExMzE4/OC04NDczLmpwZWc.jpeg",
			Description: "Viagra Boys are a Swedish post-punk band from Stockholm. The band formed in 2015, with several members coming from the bands Les Big Byrd, Pig Eyes and Nitad. In 2018, they released their debut album Street Worms. Nils Hansson, a journalist at the newspaper Dagens Nyheter described the band favorably, praising their musical style, as well as their use of black humor and satire, and rated the album a five out of five.",
			Name:        "Viagra Boys",
		},
		{
			Model:       gorm.Model{ID: 4},
			BirthDate:   time.Date(1997, 03, 04, 0, 0, 0, 0, time.Local),
			ImageURL:    "https://i.discogs.com/9wfmMJr2CBGIAoLakscp-TGHtGwSYi-mwBWE6TbmbQ8/rs:fit/g:sm/q:90/h:500/w:500/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9BLTEzNzg4/MC0xNDkwNTA0ODg0/LTU5NjMuanBlZw.jpeg",
			Description: "Ye (born Kanye Omari West on June 8, 1977) is an American rapper, singer, songwriter, record producer, and fashion designer. His musical career has been marked by dramatic change, spanning an eclectic range of influences. Outside of his music career, West has also had success in the fashion industry. ",
			Name:        "Kanye West",
		},
		{
			Model:       gorm.Model{ID: 5},
			BirthDate:   time.Date(1994, 03, 04, 0, 0, 0, 0, time.Local),
			ImageURL:    "https://i.discogs.com/t0UnVOXxYOVg6rZ7b3_j8_CmtFseOWoqyU48C-YHo7s/rs:fit/g:sm/q:90/h:426/w:500/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9BLTEwODcy/Mi0xMjgxNDMyMjMx/LmpwZWc.jpeg",
			Description: "Limp Bizkit formed in 1994 under the name \"Limp Biscut\". In 1996 they changed the spelling of the name to \"Limp Bizkut\", and again later that year to \"Limp Bizkit\" (though some concert flyers would continue to credit them as \"Limp Bizkut\" well into late 1997)",
			Name:        "Limp Bizkit",
		},
		{
			Model:       gorm.Model{ID: 6},
			BirthDate:   time.Date(1987, 03, 04, 0, 0, 0, 0, time.Local),
			ImageURL:    "https://i.discogs.com/_4U1XSCLu0s9uUKcCSS77CwchKIu1cjKspUAOHdBi7Q/rs:fit/g:sm/q:90/h:586/w:600/czM6Ly9kaXNjb2dz/LWRhdGFiYXNlLWlt/YWdlcy9BLTEyNTI0/Ni0xNTY2MzEzMjUx/LTUwMTIuanBlZw.jpeg",
			Description: "Rock band from Aberdeen, Washington, USA, formed in 1987. Emerging from the Seattle grunge scene of the late 1980s/early 1990s, Nirvana was a power trio of musicians who brought a new aesthetic to the rock scene of the time. They had already released their debut LP \"Bleach\" with Sub Pop, but their 1991 major-label debut for DGC/Geffen Records, \"Nevermind\" broke the band and grunge into the mainstream of America. Singer/guitarist Kurt Cobain's death by suicide in April 1994 brought the band to an end. ",
			Name:        "Nirvana",
		},
	})
}
