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
	SelectNewSongs(request *dtos.Paginate) (*[]dtos.SelectSongDTO, error)
	SelectSongOfGenre(request *dtos.Paginate, genreId uint) (*[]dtos.SelectSongDTO, error)
	SearchSongs(request *dtos.Paginate, keyword string) (*[]dtos.SelectSongDTO, error)
	GetSongOfArtist(artistId int) (*[]dtos.SongsOfArtistResponse, error) 
	GetSongOfAlbum(albumId int) (*[]dtos.SongsOfAlbumsResponse, error) 
	GetSongOfPlaylist(playlistId int) (*[]dtos.SelectSongDTO, error) 
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

func (service *songService) SelectNewSongs(request *dtos.Paginate) (*[]dtos.SelectSongDTO, error) {
	return service.repository.SelectNewSong(*request)
}

func (service *songService) SelectSongOfGenre(request *dtos.Paginate, genreId uint) (*[]dtos.SelectSongDTO, error) {
	return service.repository.SelectSongsOfGenre(*request, genreId)
}

func (service *songService) SearchSongs(request *dtos.Paginate, keyword string) (*[]dtos.SelectSongDTO, error) {
	return service.repository.SearchSongs(*request, keyword)
}

func (service *songService) GetSongOfArtist(artistId int) (*[]dtos.SongsOfArtistResponse, error) {
	return service.repository.GetSongsOfArtist(artistId)
}

func (service *songService) GetSongOfAlbum(albumId int) (*[]dtos.SongsOfAlbumsResponse, error) {
	return service.repository.GetSongsOfAlbums(albumId)
}

func (service *songService) GetSongOfPlaylist(playlistId int) (*[]dtos.SelectSongDTO, error) {
	return service.repository.GetSongsOfPlaylist(playlistId)
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
