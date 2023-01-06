package routes

import (
	"github.com/doduykhang/muse/pkg/constants"
	"github.com/go-chi/chi/v5"
)


func Route(r chi.Router) {
	r.Route(constants.API_ENDPOINT, func (r chi.Router) {
		r.Route(constants.AUTH_ENDPOINT, authRoute)
		r.Route("/song", songRoute)
		r.Route("/playlist", playlistRoute)
	})
}
