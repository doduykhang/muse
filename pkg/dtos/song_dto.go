package dtos

import "time"

type CreateSong struct {
	Title       string    `json:"title"`
	Duration    uint      `json:"duration"`
	ReleaseDate time.Time `json:"releaseDate"`
	Image       string    `json:"image"`
}
type UpdateSong struct {
	Title       string
	Duration    uint
	ReleaseDate time.Time
}
type SongDTO struct {
	ID          uint
	Title       string
	Duration    uint
	ReleaseDate time.Time
	Image       string
}

type SelectSongDTO struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Duration int `json:"duration"`
	ReleaseDate time.Time `json:"releaseDate"`
	Image       string `json:"image"`
	Audio string `json:"audio"`
	Artists []ArtistDTO `json:"aritsts" gorm:"serializer:json"`
	Albums []AlbumDTO `json:"albums" gorm:"serializer:json"`
}
