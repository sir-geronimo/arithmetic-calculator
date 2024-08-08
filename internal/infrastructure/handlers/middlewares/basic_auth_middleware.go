package middlewares

import (
	"context"
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	usecases "github.com/sir-geronimo/arithmetic-calculator/internal/domain/auth/use_cases"
	"github.com/sir-geronimo/arithmetic-calculator/internal/infrastructure/handlers"
	"github.com/sir-geronimo/arithmetic-calculator/internal/infrastructure/persistence"
)

type UserKey string
type HandlerFunc func(context.Context, events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)

func BasicAuthMiddleware(next HandlerFunc) HandlerFunc {

	return HandlerFunc(func(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		// Validate header is present and is not empty
		auth := strings.Split(req.Headers["authorization"], "Basic")[1]
		if auth == "" {
			return invalidAuthHeaderResponse("invalid authorization header")
		}

		// Extract credentials
		data, err := base64.StdEncoding.DecodeString(auth)
		if err != nil {
			return invalidAuthHeaderResponse("invalid authorization header")
		}

		credentials := strings.Split(string(data), ":")
		username, password := credentials[0], credentials[1]

		// Validate credentials are valid
		usecase := usecases.NewLoginUseCase(persistence.GetConnection())

		user, err := usecase.Execute(username, password)
		if err != nil {
			return invalidAuthHeaderResponse(err.Error())
		}

		// Assign `user_id` to context and continue
		ctx = context.WithValue(ctx, UserKey("user_id"), user.ID)

		return next(ctx, req)
	})
}

func invalidAuthHeaderResponse(body string) (events.APIGatewayProxyResponse, error) {
	res, err := handlers.BuildResponse(body, http.StatusUnauthorized)
	res.Headers = map[string]string{
		"WWW-Authenticate": "Basic realm=\"arithmetic-calculator\"",
	}

	return res, err
}
