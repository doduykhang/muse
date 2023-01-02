package models

type SongsAlbumsID struct {
	SongID  uint `gorm:"primarykey"`
	AlbumID uint `gorm:"primarykey"`
}

type SongsAlbums struct {
	SongsAlbumsID
}
