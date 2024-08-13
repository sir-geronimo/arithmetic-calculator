package v1

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/sir-geronimo/arithmetic-calculator/internal/app/handlers"
	"github.com/sir-geronimo/arithmetic-calculator/internal/app/handlers/middlewares"
	usecases "github.com/sir-geronimo/arithmetic-calculator/internal/app/use_cases"
	"github.com/sir-geronimo/arithmetic-calculator/internal/infrastructure/persistence"
)

func CreateOperation(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Name string
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

	userID := r.Context().Value(middlewares.UserKey).(uuid.UUID)

	usecase := usecases.NewCreateOperationUseCase(persistence.GetConnection())

	operation, err := usecase.Execute(userID, payload.Name)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if errors.Is(err, usecases.ErrInvalidOperationType) {
			statusCode = http.StatusBadRequest
		} else if errors.Is(err, usecases.ErrInsufficientBalance) {
			statusCode = http.StatusForbidden
		}

		_ = render.Render(w, r, &handlers.ErrResponse{
			Err:        err,
			StatusCode: statusCode,
			Message:    err.Error(),
		})
		return
	}

	render.JSON(w, r, operation)
}
