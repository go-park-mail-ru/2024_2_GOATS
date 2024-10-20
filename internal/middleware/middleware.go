package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

func AccessLogMiddleware(logger *zerolog.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			logger.Info().
				Str("method", r.Method).
				Str("remote_addr", r.RemoteAddr).
				Str("url", r.URL.Path).
				Dur("work_time", time.Since(start)).
				Msg("accessLogMiddleware: END")

			next.ServeHTTP(w, r)
		})
	}
}

func CorsMiddleware(logger *zerolog.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", viper.GetString("ALLOWED_ORIGIN"))
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

			if r.Method == http.MethodOptions {
				logger.Info().Msg("corsMiddleware: PREFLIGHT END")
				w.WriteHeader(http.StatusNoContent)
				return
			}

			if r.Method == http.MethodGet || r.Method == http.MethodPost {
				w.Header().Set("Content-Type", "application/json")
			}

			logger.Info().Msg("corsMiddleware: END")

			next.ServeHTTP(w, r)
		})
	}
}

func PanicMiddleware(logger *zerolog.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			defer func() {
				if err := recover(); err != nil {
					logger.Error().Msg(fmt.Sprintf("panicMiddleware: Fail to recover err: %v", err))
					http.Error(w, "Internal server error", 500)
				}
			}()

			logger.Info().Msg("panicMiddleware: END")

			next.ServeHTTP(w, r)
		})
	}
}
