package repositories

import "github.com/doduykhang/muse/pkg/models"

type SongsArtistsRepository interface {
	CrudRepository[models.SongsArtists, models.SongsArtistsID]
}

type songsArtistsRepository struct {
	CrudRepositoryImpl[models.SongsArtists, models.SongsArtistsID]
}

func NewSongsAritstsRepository() SongsArtistsRepository {
	return &songsArtistsRepository{}
}
