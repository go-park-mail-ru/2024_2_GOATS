package delivery

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api/converter"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

var _ handlers.MovieImplementationInterface = (*MovieHandler)(nil)

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
		errMsg := fmt.Errorf("getMovie action: Bad request - %w", err)
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
		errMsg := fmt.Errorf("getActor action: Bad request - %w", err)
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
