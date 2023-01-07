package service

import (
	"github.com/doduykhang/muse/pkg/dtos"
	"github.com/doduykhang/muse/pkg/repositories"
)

type AlbumService interface {
	GetNewAlbums(paginate dtos.Paginate) ([]dtos.SelectNewAlbumReponse, error)
}

type albumServiceImpl struct {
	repository repositories.AlbumRepository
}

func (service *albumServiceImpl) GetNewAlbums(paginate dtos.Paginate) ([]dtos.SelectNewAlbumReponse, error) {
	return service.repository.SelectNewAlbum(paginate)
}

func NewAlbumService() AlbumService {
	return &albumServiceImpl{
		repository: appRepositores.AlbumRepository,
	}
}
