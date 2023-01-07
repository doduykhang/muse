package service

import (
	"github.com/doduykhang/muse/pkg/dtos"
	"github.com/doduykhang/muse/pkg/models"
	"github.com/doduykhang/muse/pkg/repositories"
)

type ArtistService interface {
	ReadService[dtos.ArtistDTO, dtos.BaseID]
	SearchArtists(paginate *dtos.Paginate, keyword string) (*[]dtos.SeachArtistsResponse, error)
}

type artistServiceImpl struct {
	ReadServiceImpl[models.Aritst, models.BaseID, dtos.BaseID, dtos.ArtistDTO]
	repository repositories.ArtistRepository	
}

func (service *artistServiceImpl) SearchArtists(paginate *dtos.Paginate, keyword string) (*[]dtos.SeachArtistsResponse, error) {
	return service.repository.SearchArtist(*paginate, keyword)
}

func NewArtistService() ArtistService {
	artistService := &artistServiceImpl{
		repository: appRepositores.AritstRepository,
	}
	artistService.ReadServiceImpl.repository = appRepositores.AritstRepository
	return artistService
}
