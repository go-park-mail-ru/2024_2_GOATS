package middleware

import (
	"context"
	"crypto/rand"
	"net/http"

	"github.com/rs/zerolog/log"
)

type ctxKey int

const (
	requestIDKey ctxKey = iota
)

var (
	letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
)

func WithLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	output := make([]byte, 8)
	_, err := rand.Read(output)
	if err != nil {
		return ""
	}

	for pos := range output {
		output[pos] = letters[uint8(output[pos])%uint8(len(letters))]
	}

	return string(output)
}
