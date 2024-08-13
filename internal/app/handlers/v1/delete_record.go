package v1

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/sir-geronimo/arithmetic-calculator/internal/app/handlers"
	usecases "github.com/sir-geronimo/arithmetic-calculator/internal/app/use_cases"
	"github.com/sir-geronimo/arithmetic-calculator/internal/infrastructure/persistence"
)

func DeleteRecord(w http.ResponseWriter, r *http.Request) {
	recordID, err := uuid.Parse(chi.URLParam(r, "record_id"))
	if err != nil {
		_ = render.Render(w, r, &handlers.ErrResponse{
			Err:        err,
			StatusCode: http.StatusBadRequest,
			Message:    "invalid record_id",
		})
		return
	}

	usecase := usecases.NewDeleteRecordUseCase(persistence.GetConnection())

	record, err := usecase.Execute(recordID)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if errors.Is(err, usecases.ErrRecordNotFound) {
			statusCode = http.StatusNotFound
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
