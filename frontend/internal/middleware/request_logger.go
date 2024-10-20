package middleware

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			middleware.Logger(next).ServeHTTP(w, r)
		})
}
