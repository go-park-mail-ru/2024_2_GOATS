package middleware

import (
	"bytes"
	"context"
	"crypto/subtle"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/sessions"
	"github.com/spf13/viper"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc/metadata"

	"github.com/microcosm-cc/bluemonday"
	"github.com/rs/zerolog/log"
)

var (
	HTTPRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "path", "status"},
	)

	HTTPRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Histogram of response time for handler in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path"},
	)
)

func init() {
	prometheus.MustRegister(HTTPRequestsTotal)
	prometheus.MustRegister(HTTPRequestDuration)
}

// Обёртка для ResponseWriter, которая сохраняет статус
type statusRecorder struct {
	http.ResponseWriter
	Status int
}

func (rec *statusRecorder) WriteHeader(code int) {
	rec.Status = code // Сохраняем статус
	rec.ResponseWriter.WriteHeader(code)
}

func NewLoggingResponseWriter(w http.ResponseWriter) *statusRecorder {
	return &statusRecorder{w, http.StatusOK}
}

func AccessLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID := r.Header.Get("Req-ID")
		if reqID == "" {
			reqID = generateRequestID()
		}

		ctx := context.WithValue(r.Context(), requestIDKey, reqID)
		w.Header().Set("Req-ID", reqID)
		rec := NewLoggingResponseWriter(w)
		md := metadata.Pairs(
			"request_id", reqID,
		)

		ctx = metadata.NewOutgoingContext(ctx, md)
		start := time.Now()
		next.ServeHTTP(rec, r.WithContext(ctx))
		status := rec.Status
		logRequest(r, start, "accessLogMiddleware", reqID, status)

		w.Header().Get("")
	})
}

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", viper.GetString("ALLOWED_ORIGIN"))
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token, mode")
		w.Header().Set("Access-Control-Expose-Headers", "Content-Type, X-CSRF-Token")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}

func PanicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				logger := log.Ctx(r.Context())
				logger.Error().Msgf("panicMiddleware: Panic happend: %v", err)
				http.Error(w, "Internal server error", 500)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func logRequest(r *http.Request, start time.Time, msg string, requestID string, status int) {
	var bodyCopy bytes.Buffer
	duration := time.Since(start)

	tee := io.TeeReader(r.Body, &bodyCopy)
	r.Body = io.NopCloser(&bodyCopy)
	bodyBytes, err := io.ReadAll(tee)
	if err != nil {
		log.Error().Err(err).Msg("invalid-request-body")
	}

	log.Info().
		Str("method", r.Method).
		Str("remote_addr", r.RemoteAddr).
		Str("url", r.URL.Path).
		Str("request-id", requestID).
		Bytes("body", bodyBytes).
		Dur("work_time", duration).
		Int("status", status).
		Msg(msg)

	HTTPRequestsTotal.WithLabelValues(r.Method, r.URL.Path, strconv.Itoa(status)).Inc()
	HTTPRequestDuration.WithLabelValues(r.Method, r.URL.Path).Observe(float64(duration))
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

var store = sessions.NewCookieStore([]byte("secret-key"))

// CsrfMiddleware проверяет CSRF токен из сессии и заголовка запроса
func CsrfMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet || r.URL.Path == "/api/csrf-token" {
			next.ServeHTTP(w, r)
			return
		}

		session, err := store.Get(r, "session-name")
		if err != nil {
			http.Error(w, "Failed to get session", http.StatusInternalServerError)
			return
		}

		token, ok := session.Values["csrf_token"].(string)

		if !ok {
			http.Error(w, "Forbidden - CSRF token missing", http.StatusForbidden)
			return
		}

		csrfHeaderToken := r.Header.Get("X-CSRF-Token")
		if subtle.ConstantTimeCompare([]byte(csrfHeaderToken), []byte(token)) != 1 {
			http.Error(w, "Forbidden - CSRF token invalid", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
