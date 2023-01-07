package dtos

type GenreDTO struct {
	AuditDTO
	Name string `json:"name"`
}
