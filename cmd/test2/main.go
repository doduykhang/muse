package main

import (
	"net/http"

	"github.com/doduykhang/muse/pkg/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
    		AllowedOrigins:   []string{"https://*", "http://*"},
    		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
    		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
    		ExposedHeaders:   []string{"Link"},
    		AllowCredentials: true,
    		MaxAge:           300, // Maximum value not ignored by any of major browsers
  	}))
	r.Use(middleware.Logger)
	r.Route("/", routes.Route)
	http.ListenAndServe(":8080", r)
}
