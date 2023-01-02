package repositories

import (
	"github.com/doduykhang/muse/pkg/models"
)

type SongsAlbumsRepository interface {
	CrudRepository[models.SongsAlbums, models.SongsAlbumsID]
}

type songsAlbumsRepository struct {
	CrudRepositoryImpl[models.SongsAlbums, models.SongsAlbumsID]
}

func NewSongsAlbumsRepository() SongsAlbumsRepository {
	return &songsAlbumsRepository{}
}
