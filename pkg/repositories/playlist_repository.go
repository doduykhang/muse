package repositories

import (
	"github.com/doduykhang/muse/pkg/models"
)

type PlaylitsRepository interface {
	CrudRepository[*models.Playlist, uint]
}

type playlistRepository struct {
	CrudRepositoryImpl[*models.Playlist, uint]
}

func NewPlaylistRepository() PlaylitsRepository {
	return &playlistRepository{}
}
