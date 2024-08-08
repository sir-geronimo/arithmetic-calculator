package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	usecases "github.com/sir-geronimo/arithmetic-calculator/internal/domain/record/use_cases"
	"github.com/sir-geronimo/arithmetic-calculator/internal/infrastructure/handlers"
	"github.com/sir-geronimo/arithmetic-calculator/internal/infrastructure/persistence"
)

var usecase *usecases.FetchRecordsUseCase

func init() {
	usecase = usecases.NewFetchRecordsUseCase(
		persistence.GetConnection(),
	)
}

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	u, ok := ctx.Value("user_id").(string)
	if ok {
		log.Fatalf("empty `user_id`")
	}

	page, _ := strconv.Atoi(request.QueryStringParameters["page"])
	perPage, _ := strconv.Atoi(request.QueryStringParameters["per_page"])
	filter := request.QueryStringParameters["q"]

	records, err := usecase.Execute(u, &usecases.FetchRecordsOptions{
		Page:    page,
		PerPage: perPage,
		Filter:  filter,
	})
	if err != nil {
		return handlers.BuildResponse(err.Error(), http.StatusInternalServerError)
	}

	b, err := json.Marshal(records)
	if err != nil {
		return handlers.BuildResponse(err.Error(), http.StatusInternalServerError)
	}

	return handlers.BuildResponse(string(b), http.StatusOK)
}

func main() {
	lambda.Start(HandleRequest)
}
