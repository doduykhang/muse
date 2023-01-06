package controllers

import (
	"fmt"
	"net/http"

	"github.com/doduykhang/muse/pkg/constants"
	"github.com/doduykhang/muse/pkg/dtos"
	"github.com/doduykhang/muse/pkg/service"
	"github.com/doduykhang/muse/pkg/utils"
)

type playlistController struct {
	playlistService service.PlaylistService
}

func (controller *playlistController) GetUserPlaylists(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	paginate, ok := ctx.Value(constants.PAGINATE_DTO).(*dtos.Paginate) 	
	if !ok {
		utils.ErrorResponse(&w, constants.ERROR_GETTING_REQUEST, http.StatusBadRequest)
		return
	}

	fmt.Println("got paginate")

	userId, ok := ctx.Value(constants.ID_DTO).(uint) 	

	if !ok {
		utils.ErrorResponse(&w, constants.ERROR_GETTING_REQUEST, http.StatusBadRequest)
		return
	}

	result, err := controller.playlistService.GetUserPlayList(*paginate, userId)
	if err != nil {
		utils.ErrorResponse(&w, err.Error(), http.StatusBadRequest)
		return
	}
	utils.JsonResponse(&w, result)
}

func NewPlaylistController() playlistController {
	return playlistController{
		playlistService: businessServices.PlaylistService,
	}	
}
