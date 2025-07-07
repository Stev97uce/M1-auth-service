package auth

import (
	"net/http"

	"auth-service/internal/session"
)

func AuthMiddleware(sessionStore *session.RedisClient) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("session_token")
			if err != nil || cookie.Value == "" {
				http.Error(w, "Unauthorized - No token", http.StatusUnauthorized)
				return
			}

			_, err = sessionStore.ValidateSession(cookie.Value)
			if err != nil {
				http.Error(w, "Unauthorized - Invalid token", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
