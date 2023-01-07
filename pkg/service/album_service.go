package service

import (
	"github.com/doduykhang/muse/pkg/dtos"
	"github.com/doduykhang/muse/pkg/models"
	"github.com/doduykhang/muse/pkg/repositories"
)

type AlbumService interface {	
	ReadService[dtos.AlbumDTO, dtos.BaseID]
	GetNewAlbums(paginate dtos.Paginate) ([]dtos.SelectNewAlbumReponse, error)
	SearchAlbums(paginate dtos.Paginate, keyword string) (*[]dtos.SeachAlbumResponse, error)
	GetAlbumsOfArtist(artistId int) (*[]dtos.AlbumOfArtistResponse, error)
}

type albumServiceImpl struct {
	ReadServiceImpl[models.Album, models.BaseID, dtos.BaseID, dtos.AlbumDTO]
	repository repositories.AlbumRepository
}

func (service *albumServiceImpl) GetNewAlbums(paginate dtos.Paginate) ([]dtos.SelectNewAlbumReponse, error) {
	return service.repository.SelectNewAlbum(paginate)
}

func (service *albumServiceImpl) SearchAlbums(paginate dtos.Paginate, keyword string) (*[]dtos.SeachAlbumResponse, error) {
	return service.repository.SearchAlbums(paginate, keyword)
}

func (service *albumServiceImpl) GetAlbumsOfArtist(artistId int) (*[]dtos.AlbumOfArtistResponse, error) {
	return service.repository.GetAlbumsOfArtist(artistId)
}

func NewAlbumService() AlbumService {
	albumService := &albumServiceImpl{
		repository: appRepositores.AlbumRepository,
	}
	albumService.ReadServiceImpl.repository = appRepositores.AlbumRepository
	return albumService
}
