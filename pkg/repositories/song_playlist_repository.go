package repositories

import "github.com/doduykhang/muse/pkg/models"

type SongsPlaylitsRepository interface {
	CrudRepository[models.SongsPlaylists, models.SongsPlaylistID]
}

type songsPlaylistsRepository struct {
	CrudRepositoryImpl[models.SongsPlaylists, models.SongsPlaylistID]
}

func NewSongsPlaylistsRepository() SongsPlaylitsRepository {
	return &songsPlaylistsRepository{}
}
