package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	usecases "github.com/sir-geronimo/arithmetic-calculator/internal/domain/operation/use_cases"
	"github.com/sir-geronimo/arithmetic-calculator/internal/infrastructure/handlers"
	"github.com/sir-geronimo/arithmetic-calculator/internal/infrastructure/persistence"
)

var usecase *usecases.CreateOperationUseCase

func init() {
	usecase = usecases.NewCreateOperationUseCase(
		persistence.GetConnection(),
	)
}

var payload struct {
	Name string
}

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	json.Unmarshal([]byte(request.Body), &payload)
	if payload.Name == "" {
		return handlers.BuildResponse("invalid operation name", http.StatusBadRequest)
	}

	operation, err := usecase.Execute(payload.Name)
	if err != nil {
		return handlers.BuildResponse(err.Error(), http.StatusInternalServerError)
	}

	b, err := json.Marshal(operation)
	if err != nil {
		return handlers.BuildResponse(err.Error(), http.StatusInternalServerError)
	}

	return handlers.BuildResponse(string(b), http.StatusCreated)
}

func main() {
	lambda.Start(HandleRequest)
}
