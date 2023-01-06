package service

import (
	"github.com/doduykhang/muse/pkg/dtos"
	"github.com/doduykhang/muse/pkg/repositories"
)

type PlaylistService interface {
	GetUserPlayList(paginate dtos.Paginate, userId uint) ([]dtos.PlaylistDTO, error)
}

type playlistServiceImpl struct {
	playlistRepository repositories.PlaylitsRepository
}


func (service *playlistServiceImpl) GetUserPlayList(paginate dtos.Paginate, userId uint) ([]dtos.PlaylistDTO, error) {
	return service.playlistRepository.GetMyPlaylist(paginate, userId)
}

func NewPlayListService() PlaylistService {
	return &playlistServiceImpl{
		playlistRepository: appRepositores.PlaylistRepository,
	}
}
