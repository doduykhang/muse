package routes

import (
	"github.com/doduykhang/muse/pkg/constants"
	"github.com/doduykhang/muse/pkg/dtos"
	"github.com/doduykhang/muse/pkg/utils"
	"github.com/go-chi/chi/v5"
)

func albumRoute(r chi.Router) {
	r.With(
		utils.GetRequestQuery[dtos.Paginate](constants.PAGINATE_DTO),
	).
	Get("/new-albums", appControllers.AlbumController.GetNewAlbums)	
}
