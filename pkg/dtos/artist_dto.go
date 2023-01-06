package dtos

import "time"

type ArtistDTO struct {
	ID uint `json:"id"`
	Name string `json:"name"`
}

type SeachArtistsResponse struct {
	AuditDTO
	ID uint `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Image string `json:"image"`
	Birthdate time.Time `json:"birthdate"`
	Gender bool `json:"gender"`
}
