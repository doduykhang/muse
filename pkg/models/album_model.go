package models

type Album struct {
	BaseModel
	Name        string
	Image       string
	ReleaseDate string
	AlbumTypeID string
}
