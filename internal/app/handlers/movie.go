package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api"
	"github.com/labstack/gommon/log"
)

type MovieHandler struct {
	ApiLayer *api.Implementation
}

func NewMovieHandler(api *api.Implementation) *MovieHandler {
	return &MovieHandler{
		ApiLayer: api,
	}
}

func (m *MovieHandler) GetCollections(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		collectionsResp, errResp := m.ApiLayer.GetCollection(m.ApiLayer.Ctx, r.URL.Query())
		if errResp != nil {
			w.WriteHeader(errResp.StatusCode)
			err := json.NewEncoder(w).Encode(errResp)
			if err != nil {
				log.Errorf("error while encoding bad movie_collections response: %w", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			return
		}

		err := json.NewEncoder(w).Encode(collectionsResp)
		if err != nil {
			log.Errorf("error while encoding good movie_collections response: %w", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}
