package dtos

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
