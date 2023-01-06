package controllers

type Controllers struct {
	PlaylistController playlistController 
}

var (
	appControllers *Controllers		
)

func init() {
	appControllers = &Controllers{
		PlaylistController: NewPlaylistController(),
	}	
}

func GetControllers () *Controllers {
	return appControllers
}
