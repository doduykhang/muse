package service

import (
	"github.com/doduykhang/muse/pkg/dtos"
	"github.com/doduykhang/muse/pkg/models"
	"github.com/doduykhang/muse/pkg/repositories"
)

type PlaylistService interface {
	ReadService[dtos.PlaylistDTO, dtos.BaseID]
	GetUserPlayList(paginate dtos.Paginate, userId uint) ([]dtos.PlaylistDTO, error)
}

type playlistServiceImpl struct {
	ReadServiceImpl[models.Playlist, models.BaseID, dtos.BaseID, dtos.PlaylistDTO]
	playlistRepository repositories.PlaylitsRepository
}


func (service *playlistServiceImpl) GetUserPlayList(paginate dtos.Paginate, userId uint) ([]dtos.PlaylistDTO, error) {
	return service.playlistRepository.GetMyPlaylist(paginate, userId)
}

func NewPlayListService() PlaylistService {
	playlistService := &playlistServiceImpl{
		playlistRepository: appRepositores.PlaylistRepository,
	}
	playlistService.ReadServiceImpl.repository = appRepositores.PlaylistRepository
	return playlistService
}
