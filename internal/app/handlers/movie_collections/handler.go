package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api"
)

type MovieHandler struct {
	Context  context.Context
	ApiLayer *api.Implementation
}

func NewHandler(ctx context.Context, api *api.Implementation) *MovieHandler {
	return &MovieHandler{
		Context: ctx,
		ApiLayer: api,
	}
}

func (m *MovieHandler) GetCollections(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		collectionsResp, errResp := m.ApiLayer.GetCollection(m.Context, r.URL.Query())
		if errResp != nil {
			w.WriteHeader(errResp.StatusCode)
			json.NewEncoder(w).Encode(errResp)
			return
		}

		json.NewEncoder(w).Encode(collectionsResp)
	})
}
