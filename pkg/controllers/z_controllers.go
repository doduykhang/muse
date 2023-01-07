package controllers

type Controllers struct {
	PlaylistController *playlistController 
	SongController *SongController
	AlbumController *AlbumController
	GenreController *genreController	
}

var (
	appControllers *Controllers		
)

func init() {
	appControllers = &Controllers{
		PlaylistController: NewPlaylistController(),
		SongController: GetSongController(),	
		AlbumController: NewAlbumController(),
		GenreController: NewGenreController(),
	}	
}

func GetControllers () *Controllers {
	return appControllers
}
