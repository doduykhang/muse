package models

type Playlist struct {
	BaseModel
	Name   string
	UserID uint
}
