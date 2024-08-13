package middlewares

import (
	"context"
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/render"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/sir-geronimo/arithmetic-calculator/internal/app/handlers"
)

type UserID string

const UserKey = UserID("user_id")

func BearerTokenAuthenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			_ = render.Render(w, r, handlers.ErrUnauthorized)
			return
		}

		t := strings.Split(auth, " ")[1]
		if t == "" {
			_ = render.Render(w, r, handlers.ErrUnauthorized)
			return
		}

		token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				err := errors.New("unexpected signing method")

				return nil, err
			}

			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil || !token.Valid {
			_ = render.Render(w, r, handlers.ErrUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			_ = render.Render(w, r, handlers.ErrUnauthorized)
			return
		}

		ctx := r.Context()

		userID, err := uuid.Parse(claims["user_id"].(string))
		if err == nil {
			ctx = context.WithValue(ctx, UserKey, userID)
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
