package delivery

import (
	"context"
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

func NewMovieHandler(ctx context.Context, srv MovieServiceInterface) handlers.MovieHandlerInterface {
	return &MovieHandler{
		movieService: srv,
	}
}

func (m *MovieHandler) GetCollections(w http.ResponseWriter, r *http.Request) {
	lg := log.Ctx(r.Context())
	collectionsServResp, errServResp := m.movieService.GetCollection(r.Context())
	collectionsResp, errResp := converter.ToApiCollectionsResponse(collectionsServResp), converter.ToApiErrorResponse(errServResp)

	if errResp != nil {
		errMsg := fmt.Sprint("getCollections action: request failed - ", errResp.Errors)
		lg.Error().Msg(errMsg)
		api.Response(r.Context(), w, errResp.StatusCode, errResp)

		return
	}

	api.Response(r.Context(), w, collectionsResp.StatusCode, collectionsResp)
}

func (m *MovieHandler) GetMovie(w http.ResponseWriter, r *http.Request) {
	lg := log.Ctx(r.Context())
	mvId, err := strconv.Atoi(mux.Vars(r)["movie_id"])
	if err != nil {
		errMsg := fmt.Errorf("getMovie action: Bad request - %w", err)
		lg.Error().Msg(errMsg.Error())
		api.Response(r.Context(), w, http.StatusBadRequest, api.PreparedDefaultError("bad_request", errMsg))

		return
	}

	movieServResp, errServResp := m.movieService.GetMovie(r.Context(), mvId)
	movieResp, errResp := converter.ToApiGetMovieResponse(movieServResp), converter.ToApiErrorResponse(errServResp)

	if errResp != nil {
		errMsg := fmt.Sprint("getMovie action: request failed - ", errResp.Errors)
		lg.Error().Msg(errMsg)
		api.Response(r.Context(), w, errResp.StatusCode, errResp)

		return
	}

	api.Response(r.Context(), w, http.StatusOK, movieResp)
}

func (m *MovieHandler) GetActor(w http.ResponseWriter, r *http.Request) {
	lg := log.Ctx(r.Context())
	actorId, err := strconv.Atoi(mux.Vars(r)["actor_id"])
	if err != nil {
		errMsg := fmt.Errorf("getActor action: Bad request - %w", err)
		lg.Error().Msg(errMsg.Error())
		api.Response(r.Context(), w, http.StatusBadRequest, api.PreparedDefaultError("bad_request", errMsg))

		return
	}

	actorServResp, errServResp := m.movieService.GetActor(r.Context(), actorId)
	actorResp, errResp := converter.ToApiGetActorResponse(actorServResp), converter.ToApiErrorResponse(errServResp)

	if errResp != nil {
		errMsg := fmt.Sprint("getActor action: request failed - ", errResp.Errors)
		lg.Error().Msg(errMsg)
		api.Response(r.Context(), w, errResp.StatusCode, errResp)

		return
	}

	api.Response(r.Context(), w, http.StatusOK, actorResp)
}
