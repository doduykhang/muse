package dtos

import "time"

type AlbumDTO struct {
	ID uint `json:"id"` 
	Name string `json:"name"`
}

type AlbumOfArtistResponse struct {
	AuditDTO
	Name string `json:"name"`
	ReleaseDate time.Time `json:"releaseDate"`
	Image string `json:"image"`
	AlbumType string `json:"albumType"`
}

type SeachAlbumResponse struct {
	AuditDTO
	Name string `json:"name"`
	ReleaseDate time.Time `json:"releaseDate"`
	Image string `json:"image"`
	Artists []ArtistDTO `json:"artists"`
}

type SelectNewAlbumReponse struct {
	AuditDTO
	Name string `json:"name"`
	ReleaseDate time.Time `json:"releaseDate"`
	Image string `json:"image"`
	AlbumType string `json:"albumType"`
	Artists []ArtistDTO `json:"artists"`
}
