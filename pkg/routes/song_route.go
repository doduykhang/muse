package routes

import (
	"github.com/doduykhang/muse/pkg/controllers"
	"github.com/doduykhang/muse/pkg/dtos"
	"github.com/doduykhang/muse/pkg/utils"
	"github.com/go-chi/chi/v5"
)



func songRoute(r chi.Router) {
	r.With(
		utils.GetRequestQuery[dtos.Paginate]("request"),
	).
	Get("/", controllers.SelectSongsController)
}
