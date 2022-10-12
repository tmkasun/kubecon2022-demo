package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func HealthRoutes(r chi.Router) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		RespondwithJSON(w, 200, "healthy")
	})
}
