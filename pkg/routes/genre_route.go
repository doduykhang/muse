package routes

import "github.com/go-chi/chi/v5"

func genreRoute(r chi.Router) {
	r.Get("/", appControllers.GenreController.FindAll)
}
