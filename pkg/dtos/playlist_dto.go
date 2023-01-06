package dtos

type PlaylistDTO struct {
	AuditDTO		
	Name string `json:"name"`
	Image string `json:"image"`
}
