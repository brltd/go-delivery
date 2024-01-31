package middlewares

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/brltd/delivery/logger"
	"github.com/golang-jwt/jwt"
)

var secretKey = []byte(os.Getenv("JWT_SECRET"))

type authError struct {
	message string `string:"message"`
}

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")

		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)

			if err := json.NewEncoder(w).Encode(authError{
				message: "Unauthorized: Token not provided",
			}); err != nil {
				logger.Error("Error encoding token")
			}
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)

			if err := json.NewEncoder(w).Encode(authError{
				message: "Unauthorized",
			}); err != nil {
				logger.Error("Error encoding token")
			}
			return
		}

		if !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)

			if err := json.NewEncoder(w).Encode(authError{
				message: "Unauthorized: Invalid token",
			}); err != nil {
				logger.Error("Error encoding token")
			}
			return
		}

		next.ServeHTTP(w, r)
	})
}
