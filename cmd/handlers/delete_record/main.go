package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/google/uuid"
	usecases "github.com/sir-geronimo/arithmetic-calculator/internal/domain/record/use_cases"
	"github.com/sir-geronimo/arithmetic-calculator/internal/infrastructure/handlers"
	"github.com/sir-geronimo/arithmetic-calculator/internal/infrastructure/persistence"
)

var usecase *usecases.DeleteRecordUseCase

func init() {
	usecase = usecases.NewDeleteRecordUseCase(
		persistence.GetConnection(),
	)
}

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	recordID, err := uuid.Parse(request.PathParameters["record_id"])
	if err != nil {
		return handlers.BuildResponse("invalid `record_id` parameter", http.StatusBadRequest)
	}

	record, err := usecase.Execute(recordID)
	if err != nil {
		return handlers.BuildResponse(err.Error(), http.StatusInternalServerError)
	}

	b, err := json.Marshal(record)
	if err != nil {
		return handlers.BuildResponse(err.Error(), http.StatusInternalServerError)
	}

	return handlers.BuildResponse(string(b), http.StatusOK)
}

func main() {
	lambda.Start(HandleRequest)
}
