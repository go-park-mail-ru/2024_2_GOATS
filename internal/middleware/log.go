package middleware

import (
	"context"
	"crypto/rand"
	"net/http"
	"strings"

	"github.com/rs/zerolog/log"
)

type ctxKey int

const (
	requestIDKey ctxKey = iota
	symbols             = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
)

// WithLogger wraps logger into context
func WithLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api/room") {
			next.ServeHTTP(w, r)
			return
		}

		reqID := getRequestID(r.Context())
		logger := log.With().Str("request_id", reqID).Caller().Logger()
		ctx := logger.WithContext(r.Context())

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getRequestID(ctx context.Context) string {
	if reqID, ok := ctx.Value(requestIDKey).(string); ok {
		return reqID
	}
	return ""
}

func generateRequestID() string {
	output := make([]byte, 16)
	_, err := rand.Read(output)
	if err != nil {
		return ""
	}

	for pos := range output {
		output[pos] = symbols[uint8(output[pos])%uint8(len(symbols))]
	}

	return string(output)
}
