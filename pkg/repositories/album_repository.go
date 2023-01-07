package repositories

import (
	"github.com/doduykhang/muse/pkg/dtos"
	"github.com/doduykhang/muse/pkg/models"
)

type AlbumRepository interface {
	CrudRepository[models.Album, models.BaseID]
	GetAlbumsOfArtist(artistId int) (*[]dtos.AlbumOfArtistResponse, error)
	SearchAlbums(paginate dtos.Paginate, keyword string) (*[]dtos.SeachAlbumResponse, error)
	SelectNewAlbum(paginate dtos.Paginate) ([]dtos.SelectNewAlbumReponse, error) 
}

type albumRepository struct {
	CrudRepositoryImpl[models.Album, models.BaseID]
}

func (repositoy *albumRepository) GetAlbumsOfArtist(artistId int) (*[]dtos.AlbumOfArtistResponse, error) {
	var dtos []dtos.AlbumOfArtistResponse
	result := db.Raw("call get_albums_of_artist(?)", artistId).Scan(&dtos);
	if result.Error != nil {
		return nil, result.Error
	}
	return &dtos, nil
}

func (repositoy *albumRepository) SearchAlbums(paginate dtos.Paginate, keyword string) (*[]dtos.SeachAlbumResponse, error) {
	var dtos []dtos.SeachAlbumResponse
	result := db.Raw("call search_albums(?, ?, ?)", paginate.Size, paginate.Page, keyword).Scan(&dtos);
	if result.Error != nil {
		return nil, result.Error
	}
	return &dtos, nil
}

func (repositoy *albumRepository) SelectNewAlbum(paginate dtos.Paginate) ([]dtos.SelectNewAlbumReponse, error) {
	var dtos []dtos.SelectNewAlbumReponse
	result := db.Raw("call select_new_albums(?, ?)", paginate.Size, paginate.Page).Scan(&dtos);
	if result.Error != nil {
		return nil, result.Error
	}
	return dtos, nil
}

func NewAlbumRepository() AlbumRepository {
	return &albumRepository{}
}
