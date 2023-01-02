package repositories

import (
	"github.com/doduykhang/muse/pkg/dtos"
	"github.com/doduykhang/muse/pkg/models"
)

type SongRepository interface {
	CrudRepository[models.Song, models.BaseID]
	SelectSongs(*dtos.Paginate) (*[]dtos.SelectSongDTO, error)
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

func NewSongRepository() SongRepository {
	return &songRepository{}
}
