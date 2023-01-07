package controllers

import (
	"net/http"

	"github.com/doduykhang/muse/pkg/constants"
	"github.com/doduykhang/muse/pkg/dtos"
	"github.com/doduykhang/muse/pkg/lib"
	"github.com/doduykhang/muse/pkg/service"
	"github.com/doduykhang/muse/pkg/utils"
	"github.com/go-chi/chi/v5"
)

var (
	songController = &SongController{}
	songService    = businessServices.SongService
	
	SelectSongsController = lib.Request(songController.SelectSongs)
)

type SongController struct {
	service service.SongService
	CreateUploadController[dtos.CreateSong, dtos.SongDTO]
	UpdateController[dtos.UpdateSong, dtos.SongDTO, dtos.BaseID]
	DeleteController[dtos.BaseID]
	ReadController[dtos.SongDTO, dtos.BaseID]	
}

func (controller *SongController) SelectSongs(request *dtos.Paginate) (*[]dtos.SelectSongDTO, error) {
	dtos, err := controller.service.SelectSongs(request)
	return dtos, err
}

func (controller *SongController) SelectNewSongs(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	paginate, ok := ctx.Value(constants.PAGINATE_DTO).(*dtos.Paginate) 	
	if !ok {
		utils.ErrorResponse(&w, constants.ERROR_GETTING_REQUEST, http.StatusInternalServerError)
		return
	}

	response, err := controller.service.SelectNewSongs(paginate)
	if err != nil {
		utils.ErrorResponse(&w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.JsonResponse(&w, response)
}

func init() {
	songController.service = songService
	songController.CreateUploadController.service = songService
	songController.UpdateController.service = songService
	songController.DeleteController.service = songService
	songController.ReadController.service = songService
}

func GetSongController() *SongController {
	return songController	
}

func Test(r *chi.Mux) {

	r.With(
		utils.GetRequestForm[dtos.CreateSong],
		utils.GetFormFile("imageFile", "Image", "/image/"),
	).Post("/", songController.CreateUpload)

	r.With(
		utils.GetRequestBody[dtos.UpdateSong],
		utils.GetRequestQuery[dtos.BaseID]("id"),
	).Put("/", songController.Update)

	r.With(utils.GetRequestQuery[dtos.BaseID]("id")).Delete("/", songController.Delete)
	r.With(utils.GetRequestQuery[dtos.BaseID]("id")).Get("/", songController.FindOne)
	r.Get("/all", songController.FindAll)

	r.Post("/set-cache", func (w http.ResponseWriter, r *http.Request) {
		err := componentServices.CacheService.Set("key-muse", "value-muse")
		if err != nil {
			utils.ErrorResponse(&w, err.Error(), 500)
		}
		utils.JsonResponse(&w, "ok")
	})

	r.Get("/get-cache", func (w http.ResponseWriter, r *http.Request) {
		value, err := componentServices.CacheService.Get("key-muse")	
		if err != nil {
			utils.ErrorResponse(&w, err.Error(), 500)
		}
		utils.JsonResponse(&w, value)
	})
}


