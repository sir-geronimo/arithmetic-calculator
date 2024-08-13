package v1

import (
	"net/http"
	"strconv"

	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/sir-geronimo/arithmetic-calculator/internal/app/handlers"
	"github.com/sir-geronimo/arithmetic-calculator/internal/app/handlers/middlewares"
	usecases "github.com/sir-geronimo/arithmetic-calculator/internal/app/use_cases"
	"github.com/sir-geronimo/arithmetic-calculator/internal/infrastructure/persistence"
)

func FetchRecords(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middlewares.UserKey).(uuid.UUID)
	query := r.URL.Query()
	page, _ := strconv.Atoi(query.Get("page"))
	if page <= 0 {
		page = 1
	}
	perPage, _ := strconv.Atoi(query.Get("per_page"))
	if perPage <= 0 {
		perPage = 10
	}
	filter := query.Get("q")
	order := query.Has("order_asc")

	usecase := usecases.NewFetchRecordsUseCase(persistence.GetConnection())

	records, err := usecase.Execute(userID, &usecases.FetchRecordsOptions{
		Page:     page,
		PerPage:  perPage,
		Filter:   filter,
		OrderAsc: order,
	})
	if err != nil {
		_ = render.Render(w, r, &handlers.ErrResponse{
			Err:        err,
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}

	render.JSON(w, r, records)
}
