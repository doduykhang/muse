package repositories

import "github.com/doduykhang/muse/pkg/models"

type SongsGenresRepository interface {
	CrudRepository[models.SongsGenres, models.SongsGenresID]
}

type songsGenresRepository struct {
	CrudRepositoryImpl[models.SongsGenres, models.SongsGenresID]
}

func NewSongsGenresRepository() SongsGenresRepository {
	return &songsGenresRepository{}
}
