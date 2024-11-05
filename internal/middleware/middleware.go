package middleware

import (
	"crypto/subtle"
	"fmt"
	"net/http"
	"time"

	"github.com/microcosm-cc/bluemonday"
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
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token")

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

func sanitizeInput(input string) string {
	policy := bluemonday.UGCPolicy()
	return policy.Sanitize(input)
}

func XssMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Invalid form data", http.StatusBadRequest)
			return
		}

		for key, values := range r.Form {
			for i, v := range values {
				r.Form[key][i] = sanitizeInput(v)
			}
		}
		next.ServeHTTP(w, r)
	})
}

func CsrfMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet || r.URL.Path == "/api/csrf-token" {
			next.ServeHTTP(w, r)
			return
		}

		cookie, err := r.Cookie("csrf_token")
		if err != nil {
			http.Error(w, "Forbidden - CSRF token missing", http.StatusForbidden)
			return
		}

		csrfHeaderToken := r.Header.Get("X-CSRF-Token")
		if subtle.ConstantTimeCompare([]byte(csrfHeaderToken), []byte(cookie.Value)) != 1 {
			http.Error(w, "Forbidden - CSRF token invalid", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
