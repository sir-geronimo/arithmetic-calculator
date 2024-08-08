package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sir-geronimo/arithmetic-calculator/internal/infrastructure/handlers"
)

var payload struct {
	Username string
	Password string
}

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	json.Unmarshal([]byte(request.Body), &payload)
	if payload.Username == "" || payload.Password == "" {
		return handlers.BuildResponse("invalid payload", http.StatusBadRequest)
	}

	return handlers.BuildResponse("Login Handler", http.StatusOK)
}

func main() {
	lambda.Start(HandleRequest)
}
