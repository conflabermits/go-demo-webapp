package middleware

import (
	"net/http"

	"github.com/your-org/my-app/internal/logger"
)

func AccessLog(logger *logger.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Info("Received request",
				zap.String("method", r.Method),
				zap.String("path", r.URL.Path),
				zap.String("remoteAddr", r.RemoteAddr),
			)

			next.ServeHTTP(w, r)
		})
	}
}

func InputSanitization() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Implement input sanitization logic here

			next.ServeHTTP(w, r)
		})
	}
}
