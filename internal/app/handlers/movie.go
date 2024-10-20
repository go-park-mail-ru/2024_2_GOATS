package handlers

import (
	"net/http"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/movie/delivery"
)

type MovieHandler struct {
	ApiLayer *delivery.Implementation
	Config   *config.Config
}

func NewMovieHandler(api *delivery.Implementation, cfg *config.Config) *MovieHandler {
	return &MovieHandler{
		ApiLayer: api,
		Config:   cfg,
	}
}

func (m *MovieHandler) GetCollections(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := config.WrapContext(r.Context(), m.Config)
		collectionsResp, errResp := m.ApiLayer.GetCollection(ctx)
		if errResp != nil {
			Response(w, errResp.StatusCode, errResp)
			return
		}

		Response(w, collectionsResp.StatusCode, collectionsResp)
	})
}
