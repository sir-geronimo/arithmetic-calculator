package handlers

import (
	"log"
	"net/http"

	"github.com/go-chi/render"
)

var (
	ErrUnauthorized = &ErrResponse{StatusCode: 401, Message: "unauthorized"}
)

// ErrResponse renderer type for handling errors
type ErrResponse struct {
	// Err represent the low-level runtime error
	Err error `json:"-"`

	// StatusCode is the HttpCode used for the response
	StatusCode int `json:"code"`
	// Message is the application-level error message displayed on the response
	Message string `json:"message,omitempty"`
}

func (e *ErrResponse) Render(_ http.ResponseWriter, r *http.Request) error {
	if e.StatusCode >= http.StatusInternalServerError {
		log.Printf("\033[31m[ERROR]\033[0m %+v\n", e.Err)
	}

	render.Status(r, e.StatusCode)

	return nil
}
