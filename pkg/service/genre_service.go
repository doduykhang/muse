package service

import (
	"github.com/doduykhang/muse/pkg/dtos"
	"github.com/doduykhang/muse/pkg/models"
)

type GenreService interface {
	ReadService[dtos.GenreDTO, dtos.BaseID]
}

type genreServiceImpl struct {
	ReadServiceImpl[models.Genre, models.BaseID, dtos.BaseID, dtos.GenreDTO]
}

func NewGenreService () GenreService {
	service := &genreServiceImpl{}
	service.ReadServiceImpl.repository = appRepositores.GenreRepository
	return service
}
