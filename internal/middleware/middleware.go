package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func AccessLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		logRequest(r, start, "accessLogMiddleware: END")

		next.ServeHTTP(w, r)
	})
}

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", viper.GetString("ALLOWED_ORIGIN"))
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			log.Info().Msg("corsMiddleware: PREFLIGHT END")
			w.WriteHeader(http.StatusNoContent)
			return
		}

		if r.Method == http.MethodGet || r.Method == http.MethodPost {
			w.Header().Set("Content-Type", "application/json")
		}

		log.Info().Msg("corsMiddleware: END")

		next.ServeHTTP(w, r)
	})
}

func PanicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			if err := recover(); err != nil {
				log.Error().Msg(fmt.Sprintf("panicMiddleware: Fail to recover err: %v", err))
				http.Error(w, "Internal server error", 500)
			}
		}()

		log.Info().Msg("panicMiddleware: END")

		next.ServeHTTP(w, r)
	})
}

func logRequest(r *http.Request, start time.Time, msg string) {
	log.Info().
		Str("method", r.Method).
		Str("remote_addr", r.RemoteAddr).
		Str("url", r.URL.Path).
		Dur("work_time", time.Since(start)).
		Msg(msg)
}
