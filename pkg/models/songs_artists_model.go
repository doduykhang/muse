package models

type SongsArtistsID struct {
	SongID   uint `gorm:"primarykey"`
	ArtistID uint `gorm:"primarykey"`
}

type SongsArtists struct {
	SongsArtistsID
}
