package v1

import (
	"devops_course_app/internal/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"net/http"
)

func NewRouter(handler *chi.Mux, i usecase.InfoContract) {
	handler.Route("/info", func(r chi.Router) {
		r.Use(cors.Handler(cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Access-Control-Allow-Origin", "X-PINGOTHER", "Accept", "Authorization", "Content-Type", "X-CSRF-Token", "origin", "x-requested-with"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: true,
			MaxAge:           300,
		}))
		NewInfoRoutes(r, i)
	})

	handler.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not Found", http.StatusNotFound)
	})
}
