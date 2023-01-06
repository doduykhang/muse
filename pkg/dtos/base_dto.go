package dtos

import "time"

type AuditDTO struct {
	CreatedAt time.Time `json:"createdAt"`
	CreatedBy string `json:"createBy"`
	UpdatedAt time.Time `json:"updatedAt"`
	UpdatedBy string `json:"updatedBy"`
	DeletedAt time.Time `json:"deletedAt"`
	DeletedBy string `json:"deletedBy"`
}

type BaseID struct {
	ID uint
}

type ErrorResponse struct {
	Code    int    `json:"int"`
	Message string `json:"message"`
}

type Paginate struct {
	Size int `json:"size"`
	Page int `json:"page"`
}
