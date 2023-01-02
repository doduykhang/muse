package service

import (
	"github.com/doduykhang/muse/pkg/dtos"
	"github.com/doduykhang/muse/pkg/models"
	"github.com/doduykhang/muse/pkg/repositories"
)

type (
	SongDTO           = dtos.SongDTO
	CreateSongRequest = dtos.CreateSong
	UpdateSongRequest = dtos.UpdateSong
	BaseSongID        = dtos.BaseID
	SongModel         = models.Song
	SongID            = models.BaseID
)

var (
	songRepository = appRepositores.SongsRepository
)

type SongService interface {
	CreateService[CreateSongRequest, SongDTO]
	CreateUploadService[CreateSongRequest, SongDTO]
	UpdateService[UpdateSongRequest, SongDTO, BaseSongID]
	DeleteService[BaseSongID]
	ReadService[SongDTO, BaseSongID]
	SelectSongs(*dtos.Paginate) (*[]dtos.SelectSongDTO, error)
}

type songService struct {
	repository repositories.SongRepository
	CreateServiceImpl[SongModel, SongID, CreateSongRequest, SongDTO]
	CreateUploadServiceImpl[SongModel, SongID, CreateSongRequest, SongDTO]
	UpdateSeviceImpl[SongModel, SongID, UpdateSongRequest, BaseSongID, SongDTO]
	DeleteServiceImpl[SongModel, SongID, BaseSongID]
	ReadServiceImpl[SongModel, SongID, BaseSongID, SongDTO]
}

func (service *songService) SelectSongs(request *dtos.Paginate) (*[]dtos.SelectSongDTO, error) {
	return service.repository.SelectSongs(request)
}

func NewSongService() SongService {
	service := &songService{
		repository: songRepository,
	}
	service.CreateServiceImpl.repository = songRepository
	service.UpdateSeviceImpl.repository = songRepository
	service.DeleteServiceImpl.repository = songRepository
	service.ReadServiceImpl.repository = songRepository
	service.CreateUploadServiceImpl.repository = songRepository
	service.CreateUploadServiceImpl.fileService = NewFileService()
	return service
}
