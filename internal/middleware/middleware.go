package middleware

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func AccessLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		reqID := r.Header.Get("X-Req-ID")
		if reqID == "" {
			reqID = generateRequestID()
		}

		ctx := context.WithValue(r.Context(), requestIDKey, reqID)
		w.Header().Set("X-Req-ID", reqID)
		logRequest(r, start, "accessLogMiddleware", reqID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", viper.GetString("ALLOWED_ORIGIN"))
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		if r.Method == http.MethodGet || r.Method == http.MethodPost {
			w.Header().Set("Content-Type", "application/json")
		}

		next.ServeHTTP(w, r)
	})
}

func PanicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				lg := log.Ctx(r.Context())
				lg.Error().Msg(fmt.Sprintf("panicMiddleware: Panic happend: %v", err))
				http.Error(w, "Internal server error", 500)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func logRequest(r *http.Request, start time.Time, msg string, requestId string) {
	log.Info().
		Str("method", r.Method).
		Str("remote_addr", r.RemoteAddr).
		Str("url", r.URL.Path).
		Str("request-id", requestId).
		Dur("work_time", time.Since(start)).
		Msg(msg)
}
