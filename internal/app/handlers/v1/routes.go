package v1

import (
	"github.com/go-chi/chi"
	"github.com/sir-geronimo/arithmetic-calculator/internal/app/handlers/middlewares"
)

func Routes(r chi.Router) {
	r.Route("/v1", func(r chi.Router) {
		r.Post("/login", Login)

		r.Group(func(r chi.Router) {
			r.Use(middlewares.BearerTokenAuthenticator)

			r.Post("/operations", CreateOperation)
			r.Post("/operations/{operation_id}/perform", PerformOperation)
			r.Get("/records", FetchRecords)
			r.Delete("/records/{record_id}", DeleteRecord)
		})
	})
}
