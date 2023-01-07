package repositories

import (
	"github.com/doduykhang/muse/pkg/dtos"
	"github.com/doduykhang/muse/pkg/models"
)

type SongRepository interface {
	CrudRepository[models.Song, models.BaseID]
	SelectSongs(*dtos.Paginate) (*[]dtos.SelectSongDTO, error)
	GetSongsOfAlbums(albumId int) (*[]dtos.SongsOfAlbumsResponse, error)
	GetSongsOfArtist(artistId int) (*[]dtos.SongsOfArtistResponse, error)
	GetSongsOfPlaylist(playlistId int) (*[]dtos.SelectSongDTO, error)
	SearchSongs(paginate dtos.Paginate, keyword string) (*[]dtos.SelectSongDTO, error) 
	SelectNewSong(paginate dtos.Paginate) (*[]dtos.SelectSongDTO, error)
	SelectSongsOfGenre(paginate dtos.Paginate, genreId uint) (*[]dtos.SelectSongDTO, error)
}

type songRepository struct {
	CrudRepositoryImpl[models.Song, models.BaseID]
}

func (repositoy *songRepository) SelectSongs(request *dtos.Paginate) (*[]dtos.SelectSongDTO, error) {
	var dtos []dtos.SelectSongDTO
	result := db.Raw("call select_songs(?, ?)", request.Size, request.Page).Scan(&dtos);
	if result.Error != nil {
		return nil, result.Error
	}
	return &dtos, nil
}

func (repositoy *songRepository) GetSongsOfAlbums(albumId int) (*[]dtos.SongsOfAlbumsResponse, error) {
	var dtos []dtos.SongsOfAlbumsResponse
	result := db.Raw("call get_songs_of_albums(?)", albumId).Scan(&dtos);
	if result.Error != nil {
		return nil, result.Error
	}
	return &dtos, nil
}

func (repositoy *songRepository) GetSongsOfArtist(artistId int) (*[]dtos.SongsOfArtistResponse, error) {
	var dtos []dtos.SongsOfArtistResponse
	result := db.Raw("call get_songs_of_artist(?)", artistId).Scan(&dtos);
	if result.Error != nil {
		return nil, result.Error
	}
	return &dtos, nil
}

func (repositoy *songRepository) GetSongsOfPlaylist(playlistId int) (*[]dtos.SelectSongDTO, error) {
	var dtos []dtos.SelectSongDTO
	result := db.Raw("call get_songs_of_playlist(?)", playlistId).Scan(&dtos);
	if result.Error != nil {
		return nil, result.Error
	}
	return &dtos, nil
}

func (repositoy *songRepository) SearchSongs(paginate dtos.Paginate, keyword string) (*[]dtos.SelectSongDTO, error) {
	var dtos []dtos.SelectSongDTO
	result := db.Raw("call search_songs(?, ?, ?)", paginate.Size, paginate.Page, keyword).Scan(&dtos);
	if result.Error != nil {
		return nil, result.Error
	}
	return &dtos, nil
}

func (repositoy *songRepository) SelectNewSong(paginate dtos.Paginate) (*[]dtos.SelectSongDTO, error) {
	var dtos []dtos.SelectSongDTO
	result := db.Raw("call select_new_songs(?, ?)", paginate.Size, paginate.Page).Scan(&dtos);
	if result.Error != nil {
		return nil, result.Error
	}
	return &dtos, nil
}

func (repositoy *songRepository) SelectSongsOfGenre(paginate dtos.Paginate, genreId uint) (*[]dtos.SelectSongDTO, error) {
	var dtos []dtos.SelectSongDTO
	result := db.Raw("call select_songs_of_genre(?, ?, ?)", paginate.Size, paginate.Page, genreId).Scan(&dtos);
	if result.Error != nil {
		return nil, result.Error
	}
	return &dtos, nil
}

func NewSongRepository() SongRepository {
	return &songRepository{}
}
