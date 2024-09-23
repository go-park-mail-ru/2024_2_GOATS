package handler

import (
	"context"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api"
)

func MovieCollections(ctx context.Context, api *api.Implementation, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		api.GetCollection(ctx, r.URL.Query())
	})
}
