package repositories

type appRepositories struct {
	AccountRepository        AccountRepository
	AlbumRepository          AlbumRepository
	AritstRepository         ArtistRepository
	GenreRepository          GenreRepository
	PlaylistRepository       PlaylitsRepository
	SongsAlbumsRepository    SongsAlbumsRepository
	SongsArtistsRepository   SongsArtistsRepository
	SongsGenresRepository    SongsGenresRepository
	SongsPlaylistsRepository SongsPlaylitsRepository
	SongsRepository          SongRepository
	UserRepository           UserRepository
}

var (
	repositories = &appRepositories{
		AccountRepository:        NewAccountRepository(),
		AlbumRepository:          NewAlbumRepository(),
		AritstRepository:         NewArtistRepository(),
		GenreRepository:          NewGenereRepository(),
		PlaylistRepository:       NewPlaylistRepository(),
		SongsAlbumsRepository:    NewSongsAlbumsRepository(),
		SongsArtistsRepository:   NewSongsAritstsRepository(),
		SongsGenresRepository:    NewSongsGenresRepository(),
		SongsPlaylistsRepository: NewSongsPlaylistsRepository(),
		SongsRepository:          NewSongRepository(),
		UserRepository:           NewUserRepository(),
	}
)

func GetAppRepositories() *appRepositories {
	return repositories
}
