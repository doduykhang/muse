package repositories

import (
	"github.com/doduykhang/muse/pkg/dtos"
	"github.com/doduykhang/muse/pkg/models"
)

type PlaylitsRepository interface {
	CrudRepository[*models.Playlist, uint]
	GetMyPlaylist(paginate dtos.Paginate, userId uint) ([]dtos.PlaylistDTO, error)
}

type playlistRepository struct {
	CrudRepositoryImpl[*models.Playlist, uint]
}

func (repositoy *playlistRepository) GetMyPlaylist(paginate dtos.Paginate, userId uint) ([]dtos.PlaylistDTO, error) {
	var dtos []dtos.PlaylistDTO
	result := db.Raw("call select_my_playlist(?, ?, ?)", paginate.Size, paginate.Page, userId).Scan(&dtos);
	if result.Error != nil {
		return nil, result.Error
	}
	return dtos, nil
}

func NewPlaylistRepository() PlaylitsRepository {
	return &playlistRepository{}
}
