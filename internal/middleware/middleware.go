package middleware

import (
	"bytes"
	"context"
	"crypto/subtle"
	"io"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-park-mail-ru/2024_2_GOATS/metrics"
	"github.com/gorilla/mux"
	"google.golang.org/grpc/metadata"

	"github.com/gorilla/sessions"
	"github.com/microcosm-cc/bluemonday"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type statusRecorder struct {
	http.ResponseWriter
	Path   string
	Status int
}

func (rec *statusRecorder) WriteHeader(code int) {
	if strings.HasPrefix(rec.Path, "/api/room/join") {
		return
	}

	if rec.Status != http.StatusOK {
		return
	}

	rec.Status = code
	rec.ResponseWriter.WriteHeader(code)
}

// NewLoggingResponseWriter wraps resonseWriter with status
func NewLoggingResponseWriter(w http.ResponseWriter, path string) *statusRecorder {
	return &statusRecorder{w, path, http.StatusOK}
}

// AccessLogMiddleware logs any request
func AccessLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID := r.Header.Get("Req-ID")
		if reqID == "" {
			reqID = generateRequestID()
		}

		ctx := context.WithValue(r.Context(), requestIDKey, reqID)
		w.Header().Set("Req-ID", reqID)
		var customRec *statusRecorder
		if !strings.HasPrefix(r.URL.Path, "/api/room/join") {
			customRec = NewLoggingResponseWriter(w, r.URL.Path)
		}

		md := metadata.Pairs(
			"request_id", reqID,
		)

		ctx = metadata.NewOutgoingContext(ctx, md)
		start := time.Now()

		if customRec != nil {
			next.ServeHTTP(customRec, r.WithContext(ctx))
			status := customRec.Status
			logRequest(r, start, "accessLogMiddleware", reqID, status, requestPath(w, r))
		} else {
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}

// CorsMiddleware set CORS Headers
func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", viper.GetString("ALLOWED_ORIGIN"))
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token, mode")
		w.Header().Set("Access-Control-Expose-Headers", "Content-Type, X-CSRF-Token")

		if r.Method == http.MethodOptions {
			if rw, ok := w.(*statusRecorder); ok {
				rw.WriteHeader(http.StatusNoContent)
			} else {
				w.WriteHeader(http.StatusNoContent)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}

// PanicMiddleware prevents any panic
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

func logRequest(r *http.Request, start time.Time, msg string, requestID string, status int, path string) {
	var bodyCopy bytes.Buffer
	duration := time.Since(start)

	tee := io.TeeReader(r.Body, &bodyCopy)
	r.Body = io.NopCloser(&bodyCopy)
	bodyBytes, err := io.ReadAll(tee)
	if err != nil {
		log.Error().Err(err).Msg("invalid-request-body")
	}

	if isWebSocket(r) {
		log.Info().
			Str("websocket", path).
			Str("request-id", requestID).
			Bytes("body", bodyBytes).
			Str("real_ip", realIP(r)).
			Int64("content_length", r.ContentLength).
			Str("start_time", start.Format(time.RFC3339)).
			Str("duration_human", duration.String()).
			Int64("duration_ms", duration.Milliseconds()).
			Msg(msg)

		return
	}

	log.Info().
		Str("method", r.Method).
		Str("remote_addr", r.RemoteAddr).
		Str("url", path).
		Str("request-id", requestID).
		Bytes("body", bodyBytes).
		Dur("work_time", duration).
		Int("status", status).
		Str("user_agent", r.UserAgent()).
		Str("host", r.Host).
		Str("real_ip", realIP(r)).
		Int64("content_length", r.ContentLength).
		Str("start_time", start.Format(time.RFC3339)).
		Str("duration_human", duration.String()).
		Int64("duration_ms", duration.Milliseconds()).
		Msg(msg)

	metrics.HTTPRequestTotal.WithLabelValues(r.Method, path, strconv.Itoa(status)).Inc()
	metrics.HTTPRequestDuration.WithLabelValues(r.Method, path).Observe(duration.Seconds())
}

func sanitizeInput(input string) string {
	policy := bluemonday.UGCPolicy()
	return policy.Sanitize(input)
}

// XSSMiddleware sanitize inputs
func XSSMiddleware(next http.Handler) http.Handler {
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

// CsrfMiddleware checks CSRF-Token
func CsrfMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet || r.URL.Path == "/api/csrf-token" || strings.HasPrefix(r.URL.Path, "/api/payments") {
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

func realIP(r *http.Request) string {
	rIP := r.Header.Get("X-Real-IP")
	if rIP == "" {
		rIP = r.Header.Get("X-Forwarded-For")
		if rIP != "" {
			rIP = strings.Split(rIP, ",")[0]
		}
	}

	if rIP == "" {
		rIP, _, _ = net.SplitHostPort(r.RemoteAddr)
	}

	return rIP
}

func requestPath(w http.ResponseWriter, r *http.Request) string {
	route := mux.CurrentRoute(r)
	if route == nil {
		http.Error(w, "Route not found", http.StatusNotFound)
		return ""
	}

	return route.GetName()
}

func isWebSocket(r *http.Request) bool {
	return r.Header.Get("Upgrade") == "websocket" && r.Header.Get("Connection") == "Upgrade"
}
