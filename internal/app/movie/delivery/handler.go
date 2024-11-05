package delivery

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api/converter"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

var _ handlers.MovieHandlerInterface = (*MovieHandler)(nil)

type MovieHandler struct {
	movieService MovieServiceInterface
}

func NewMovieHandler(srv MovieServiceInterface) handlers.MovieHandlerInterface {
	return &MovieHandler{
		movieService: srv,
	}
}

func (m *MovieHandler) GetCollections(w http.ResponseWriter, r *http.Request) {
	logger := log.Ctx(r.Context())
	collectionsServResp, errServResp := m.movieService.GetCollection(r.Context())
	collectionsResp, errResp := converter.ToApiCollectionsResponse(collectionsServResp), converter.ToApiErrorResponse(errServResp)

	if errResp != nil {
		errMsg := errors.New("failed to get collections")
		logger.Error().Err(errMsg).Interface("getCollectionsResp", errResp).Msg("request_failed")
		api.Response(r.Context(), w, errResp.StatusCode, errResp)

		return
	}

	logger.Info().Interface("getCollectionsResp", collectionsResp).Msg("getCollections success")

	api.Response(r.Context(), w, collectionsResp.StatusCode, collectionsResp)
}

func (m *MovieHandler) GetMovie(w http.ResponseWriter, r *http.Request) {
	logger := log.Ctx(r.Context())
	mvID, err := strconv.Atoi(mux.Vars(r)["movie_id"])
	if err != nil {
		errMsg := fmt.Errorf("getMovie action: Bad request - %w", err)
		logger.Error().Err(errMsg).Msg("bad_request")
		api.Response(r.Context(), w, http.StatusBadRequest, api.PreparedDefaultError("bad_request", errMsg))

		return
	}

	movieServResp, errServResp := m.movieService.GetMovie(r.Context(), mvID)
	movieResp, errResp := converter.ToApiGetMovieResponse(movieServResp), converter.ToApiErrorResponse(errServResp)

	if errResp != nil {
		errMsg := errors.New("failed to get movie")
		logger.Error().Err(errMsg).Interface("getMovieResp", errResp).Msg("request_failed")
		api.Response(r.Context(), w, errResp.StatusCode, errResp)

		return
	}

	logger.Info().Interface("getMovieResp", movieResp).Msg("getMovie success")

	api.Response(r.Context(), w, http.StatusOK, movieResp)
}

func (m *MovieHandler) GetActor(w http.ResponseWriter, r *http.Request) {
	logger := log.Ctx(r.Context())
	actorID, err := strconv.Atoi(mux.Vars(r)["actor_id"])
	if err != nil {
		errMsg := fmt.Errorf("getActor action: Bad request - %w", err)
		logger.Error().Err(errMsg).Msg("bad_request")
		api.Response(r.Context(), w, http.StatusBadRequest, api.PreparedDefaultError("bad_request", errMsg))

		return
	}

	actorServResp, errServResp := m.movieService.GetActor(r.Context(), actorID)
	actorResp, errResp := converter.ToApiGetActorResponse(actorServResp), converter.ToApiErrorResponse(errServResp)

	if errResp != nil {
		errMsg := errors.New("failed to getActor")
		logger.Error().Err(errMsg).Interface("actorResp", errResp).Msg("request_failed")
		api.Response(r.Context(), w, errResp.StatusCode, errResp)

		return
	}

	logger.Info().Interface("actorResp", actorResp).Msg("getActor success")

	api.Response(r.Context(), w, http.StatusOK, actorResp)
}
