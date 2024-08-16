package v1

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/sir-geronimo/arithmetic-calculator/internal/app/handlers"
	"github.com/sir-geronimo/arithmetic-calculator/internal/app/handlers/middlewares"
	usecases "github.com/sir-geronimo/arithmetic-calculator/internal/app/use_cases"
	"github.com/sir-geronimo/arithmetic-calculator/internal/infrastructure/persistence"
)

func GetBalance(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middlewares.UserKey).(uuid.UUID)

	usecase := usecases.NewGetBalanceUseCase(persistence.GetConnection())

	balance, err := usecase.Execute(userID)
	if err != nil {
		statusCode := http.StatusInternalServerError

		_ = render.Render(w, r, &handlers.ErrResponse{
			Err:        err,
			StatusCode: statusCode,
			Message:    err.Error(),
		})
		return
	}

	render.JSON(w, r, balance)
}
