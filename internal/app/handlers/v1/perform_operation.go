package v1

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/sir-geronimo/arithmetic-calculator/internal/app/handlers"
	"github.com/sir-geronimo/arithmetic-calculator/internal/app/handlers/middlewares"
	usecases "github.com/sir-geronimo/arithmetic-calculator/internal/app/use_cases"
	"github.com/sir-geronimo/arithmetic-calculator/internal/infrastructure/persistence"
)

func PerformOperation(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Num1 string `json:"num1,omitempty"`
		Num2 string `json:"num2,omitempty"`
	}

	operationID, err := uuid.Parse(chi.URLParam(r, "operation_id"))
	if err != nil {
		_ = render.Render(w, r, &handlers.ErrResponse{
			Err:        err,
			StatusCode: http.StatusBadRequest,
			Message:    "invalid operation_id",
		})
		return
	}

	userID := r.Context().Value(middlewares.UserKey).(uuid.UUID)

	if r.Body != nil && r.ContentLength > 0 {
		err = json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			_ = render.Render(w, r, &handlers.ErrResponse{
				Err:        err,
				StatusCode: http.StatusBadRequest,
				Message:    err.Error(),
			})
			return
		}

		defer r.Body.Close()
	}

	usecase := usecases.NewPerformOperationUseCase(persistence.GetConnection())

	num1, _ := strconv.Atoi(payload.Num1)
	num2, _ := strconv.Atoi(payload.Num2)

	record, err := usecase.Execute(&usecases.PerformOperationRequest{
		OperationID: operationID,
		UserID:      userID,
		Num1:        num1,
		Num2:        num2,
	})
	if err != nil {
		statusCode := http.StatusInternalServerError
		if errors.Is(err, usecases.ErrOperationNotFound) {
			statusCode = http.StatusNotFound
		} else if errors.Is(err, usecases.ErrOperationAlreadyPerformed) ||
			errors.Is(err, usecases.ErrInvalidOperationPayload) {
			statusCode = http.StatusBadRequest
		} else if errors.Is(err, usecases.ErrUnableToFindRecord) {
			statusCode = http.StatusNotFound
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

	render.JSON(w, r, record)
}
