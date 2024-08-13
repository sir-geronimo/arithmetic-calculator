package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sir-geronimo/arithmetic-calculator/internal/app/handlers"
	v1 "github.com/sir-geronimo/arithmetic-calculator/internal/app/handlers/v1"
)

var (
	port = fmt.Sprintf(":%s", os.Getenv("PORT"))
)

func Run() error {
	r := chi.NewRouter()

	r.Use(
		middleware.Logger,
		cors.Handler(cors.Options{
			AllowedOrigins: []string{"https://*", "http://*"},
			AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
			AllowedHeaders: []string{"*"},
		}),
		middleware.AllowContentType("application/json", "text/plain"),
		middleware.CleanPath,
		middleware.Recoverer,
	)

	r.Group(v1.Routes)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		_ = render.Render(w, r, &handlers.ErrResponse{
			Err:        nil,
			StatusCode: http.StatusNotFound,
			Message:    "resource not found",
		})
	})

	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		_ = render.Render(w, r, &handlers.ErrResponse{
			Err:        nil,
			StatusCode: http.StatusMethodNotAllowed,
			Message:    "method not allowed",
		})
	})

	log.Printf("Server started at: [%s]", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err)
	}

	return nil
}
