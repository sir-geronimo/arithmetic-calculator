package v1

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"github.com/sir-geronimo/arithmetic-calculator/internal/app/handlers"
	usecases "github.com/sir-geronimo/arithmetic-calculator/internal/app/use_cases"
	"github.com/sir-geronimo/arithmetic-calculator/internal/infrastructure/persistence"
)

var (
	loginUseCase = usecases.NewLoginUseCase(persistence.GetConnection())
)

func Login(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Username string
		Password string
	}

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		_ = render.Render(w, r, &handlers.ErrResponse{
			Err:        err,
			StatusCode: http.StatusBadRequest,
			Message:    "invalid payload",
		})
		return
	}
	defer r.Body.Close()

	token, err := loginUseCase.Execute(payload.Username, payload.Password)
	if err != nil {
		statusCode := http.StatusUnauthorized
		if !errors.Is(err, usecases.ErrInvalidCredentials) {
			statusCode = http.StatusInternalServerError
		}

		_ = render.Render(w, r, &handlers.ErrResponse{
			Err:        err,
			StatusCode: statusCode,
			Message:    err.Error(),
		})
		return
	}

	render.JSON(w, r, token)
}
