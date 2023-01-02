package models

type SongsGenresID struct {
	SongID  uint `gorm:"primarykey"`
	GenreID uint `gorm:"primarykey"`
}

type SongsGenres struct {
	SongsGenresID
}
