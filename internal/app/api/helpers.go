package api

import (
	"context"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/logger"
)

func GetRequestId(w http.ResponseWriter) string {
	return w.Header().Get("request-id")
}

func BaseContext(w http.ResponseWriter, r *http.Request, lg *logger.BaseLogger) context.Context {
	ctx := context.WithValue(r.Context(), "request-id", GetRequestId(w))
	return config.WrapLoggerContext(ctx, lg)
}
