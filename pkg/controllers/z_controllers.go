package controllers

type Controllers struct {
	PlaylistController *playlistController 
	SongController *SongController
	AlbumController *AlbumController
}

var (
	appControllers *Controllers		
)

func init() {
	appControllers = &Controllers{
		PlaylistController: NewPlaylistController(),
		SongController: GetSongController(),	
		AlbumController: NewAlbumController(),
	}	
}

func GetControllers () *Controllers {
	return appControllers
}
