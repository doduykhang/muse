package repositories

import (
	"github.com/doduykhang/muse/pkg/models"
)

type AlbumRepository interface {
	CrudRepository[*models.Album, uint]
}

type albumRepository struct {
	CrudRepositoryImpl[*models.Album, uint]
}

func NewAlbumRepository() AlbumRepository {
	return &albumRepository{}
}
