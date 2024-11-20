package delivery

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api/converter"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

type MovieHandler struct {
	movieService MovieServiceInterface
	logger       *zerolog.Logger
}

func NewMovieHandler(ctx context.Context, srv MovieServiceInterface) *MovieHandler {
	return &MovieHandler{
		movieService: srv,
		logger:       &config.FromContext(ctx).Logger,
	}
}

func (m *MovieHandler) GetCollections(w http.ResponseWriter, r *http.Request) {
	ctx := m.logger.WithContext(r.Context())
	collectionsServResp, errServResp := m.movieService.GetCollection(ctx)
	collectionsResp, errResp := converter.ToApiCollectionsResponse(collectionsServResp), converter.ToApiErrorResponse(errServResp)

	if errResp != nil {
		api.Response(w, errResp.StatusCode, errResp)
		return
	}

	api.Response(w, collectionsResp.StatusCode, collectionsResp)
}

func (m *MovieHandler) GetMovie(w http.ResponseWriter, r *http.Request) {
	ctx := m.logger.WithContext(r.Context())
	mvId, err := strconv.Atoi(mux.Vars(r)["movie_id"])
	if err != nil {
		errMsg := fmt.Errorf("GetMovie action: Bad request - %w", err)
		m.logger.Error().Msg(errMsg.Error())
		api.Response(w, http.StatusBadRequest, api.PreparedDefaultError("bad_request", errMsg))

		return
	}

	movieServResp, errServResp := m.movieService.GetMovie(ctx, mvId)
	movieResp, errResp := converter.ToApiGetMovieResponse(movieServResp), converter.ToApiErrorResponse(errServResp)

	if errResp != nil {
		api.Response(w, errResp.StatusCode, errResp)
		return
	}

	api.Response(w, http.StatusOK, movieResp)
}

func (m *MovieHandler) GetActor(w http.ResponseWriter, r *http.Request) {
	ctx := m.logger.WithContext(r.Context())
	actorId, err := strconv.Atoi(mux.Vars(r)["actor_id"])
	if err != nil {
		errMsg := fmt.Errorf("GetActor action: Bad request - %w", err)
		m.logger.Error().Msg(errMsg.Error())
		api.Response(w, http.StatusBadRequest, api.PreparedDefaultError("bad_request", errMsg))

		return
	}

	actorServResp, errServResp := m.movieService.GetActor(ctx, actorId)
	actorResp, errResp := converter.ToApiGetActorResponse(actorServResp), converter.ToApiErrorResponse(errServResp)

	if errResp != nil {
		api.Response(w, errResp.StatusCode, errResp)
		return
	}

	api.Response(w, http.StatusOK, actorResp)
}

func (h *MovieHandler) SearchMovies(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	if query == "" {
		http.Error(w, "query parameter is required", http.StatusBadRequest)
		return
	}

	movies, err := h.movieService.SearchMovies(r.Context(), query)
	if err != nil {
		http.Error(w, "search error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(movies); err != nil {
		http.Error(w, "response error: "+err.Error(), http.StatusInternalServerError)
	}
}

func (h *MovieHandler) SearchActors(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	if query == "" {
		http.Error(w, "query parameter is required", http.StatusBadRequest)
		return
	}

	movies, err := h.movieService.SearchActors(r.Context(), query)
	if err != nil {
		http.Error(w, "search error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(movies); err != nil {
		http.Error(w, "response error: "+err.Error(), http.StatusInternalServerError)
	}
}
