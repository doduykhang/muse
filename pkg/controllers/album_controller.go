package controllers

import (
	"net/http"

	"github.com/doduykhang/muse/pkg/constants"
	"github.com/doduykhang/muse/pkg/dtos"
	"github.com/doduykhang/muse/pkg/service"
	"github.com/doduykhang/muse/pkg/utils"
)

type AlbumController struct {
	service service.AlbumService
}

func (controller *AlbumController) GetNewAlbums(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()		
	paginate, ok := ctx.Value(constants.PAGINATE_DTO).(*dtos.Paginate)
	if !ok {
		utils.ErrorResponse(&w, constants.ERROR_GETTING_REQUEST, http.StatusInternalServerError)
		return
	}

	response, err := controller.service.GetNewAlbums(*paginate)	
	if err != nil {
		utils.ErrorResponse(&w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.JsonResponse(&w, response)
}

func NewAlbumController () *AlbumController {
	return &AlbumController{
		service: businessServices.AlbumService,
	}
}
