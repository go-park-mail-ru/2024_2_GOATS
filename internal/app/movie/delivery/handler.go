package delivery

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api/converter"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api/handlers"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

const (
	genresFilter = "genres"
)

// MovieHandler http movie handler
type MovieHandler struct {
	movieService MovieServiceInterface
}

// NewMovieHandler returns an instance of MovieHandlerInterface
func NewMovieHandler(srv MovieServiceInterface) handlers.MovieHandlerInterface {
	return &MovieHandler{
		movieService: srv,
	}
}

// TODO раскоментить к 4му РК

// func (m *MovieHandler) GetMovieByGenre(w http.ResponseWriter, r *http.Request) {
// 	logger := log.Ctx(r.Context())
// 	genre := r.URL.Query().Get("genre")

// 	if genre == "" {
// 		errMsg := errors.New("incorrect genre was given")
// 		err := errVals.NewDeliveryError(
// 			http.StatusBadRequest,
// 			[]errVals.ErrorItem{
// 				errVals.NewErrorItem("bad_request", errVals.NewCustomError(errMsg.Error())),
// 			},
// 		)

// 		logger.Error().Err(errMsg).Interface("getMovieByGenre", err).Msg("request_failed")
// 		api.Response(r.Context(), w, err.HTTPStatus, err)

// 		return
// 	}

// 	srvResp, errServResp := m.movieService.GetMovieByGenre(r.Context(), genre)
// 	resp, errResp := converter.ToAPIMovieShortInfos(srvResp), errVals.ToDeliveryErrorFromService(errServResp)
// 	if errResp != nil {
// 		errMsg := errors.New("failed to get movies by genre")
// 		logger.Error().Err(errMsg).Interface("getMovieByGenre", errResp).Msg("request_failed")
// 		api.Response(r.Context(), w, errResp.HTTPStatus, errResp)

// 		return
// 	}

// 	logger.Info().Interface("getMovieByGenre", resp).Msg("byGenre success")

// 	api.Response(r.Context(), w, http.StatusOK, resp)
// }

// GetCollections gets movie collections handler
func (m *MovieHandler) GetCollections(w http.ResponseWriter, r *http.Request) {
	m.collectMovieData(w, r, "")
}

// GetGenres gets genres collections handler
func (m *MovieHandler) GetGenres(w http.ResponseWriter, r *http.Request) {
	m.collectMovieData(w, r, genresFilter)
}

func (m *MovieHandler) collectMovieData(w http.ResponseWriter, r *http.Request, filter string) {
	logger := log.Ctx(r.Context())

	collectionsServResp, errServResp := m.movieService.GetCollection(r.Context(), filter)
	collectionsResp, errResp := converter.ToAPICollectionsResponse(collectionsServResp), errVals.ToDeliveryErrorFromService(errServResp)

	if errResp != nil {
		errMsg := errors.New("failed to get collections")
		logger.Error().Err(errMsg).Interface("getCollectionsResp", errResp).Msg("request_failed")
		api.Response(r.Context(), w, errResp.HTTPStatus, errResp)

		return
	}

	logger.Info().Interface("getCollectionsResp", collectionsResp).Msg("getCollections success")

	api.Response(r.Context(), w, http.StatusOK, collectionsResp)
}

// GetMovie gets movie handler
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
	rating, errServResp := m.movieService.GetUserRating(r.Context(), mvID)

	movieResp, errResp := converter.ToAPIGetMovieResponse(movieServResp, int64(rating)), errVals.ToDeliveryErrorFromService(errServResp)

	if errResp != nil {
		errMsg := errors.New("failed to get movie_service")
		logger.Error().Err(errMsg).Interface("getMovieResp", errResp).Msg("request_failed")
		api.Response(r.Context(), w, errResp.HTTPStatus, errResp)

		return
	}

	logger.Info().Interface("getMovieResp", movieResp).Msg("getMovie success")

	api.Response(r.Context(), w, http.StatusOK, movieResp)
}

// GetActor gets actor handler
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
	actorResp, errResp := converter.ToAPIGetActorResponse(actorServResp), errVals.ToDeliveryErrorFromService(errServResp)

	if errResp != nil {
		errMsg := errors.New("failed to getActor")
		logger.Error().Err(errMsg).Interface("actorResp", errResp).Msg("request_failed")
		api.Response(r.Context(), w, errResp.HTTPStatus, errResp)

		return
	}

	logger.Info().Interface("actorResp", actorResp).Msg("getActor success")

	api.Response(r.Context(), w, http.StatusOK, actorResp)
}

// SearchMovies search movies handler
func (m *MovieHandler) SearchMovies(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	if query == "" {
		http.Error(w, "query parameter is required", http.StatusBadRequest)
		return
	}

	movies, err := m.movieService.SearchMovies(r.Context(), query)
	if err != nil {
		http.Error(w, "search error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var movieResponses api.MovieSearchList
	if len(movies) == 0 {
		movieResponses = append(movieResponses, api.MovieSearchData{})
	}
	for _, movie := range movies {
		movieResponses = append(movieResponses, api.MovieSearchData{
			ID:          movie.ID,
			Title:       movie.Title,
			CardURL:     movie.CardURL,
			AlbumURL:    movie.AlbumURL,
			Rating:      strconv.FormatFloat(float64(movie.Rating), 'f', -1, 32),
			ReleaseDate: movie.ReleaseDate,
			MovieType:   movie.MovieType,
			Country:     movie.Country,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	jsonData, err := movieResponses.MarshalJSON()
	if err != nil {
		http.Error(w, "response error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

// SearchActors search actors handler
func (m *MovieHandler) SearchActors(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	if query == "" {
		http.Error(w, "query parameter is required", http.StatusBadRequest)
		return
	}

	actors, err := m.movieService.SearchActors(r.Context(), query)
	if err != nil {
		http.Error(w, "search error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var actorResponses api.ActorSearchList

	if len(actors) == 0 {
		actorResponses = append(actorResponses, api.ActorSearchData{})
	}

	for _, actor := range actors {
		actorResponses = append(actorResponses, api.ActorSearchData{
			ID:       actor.ID,
			FullName: actor.Name,
			PhotoURL: actor.BigPhotoURL,
			Country:  actor.Country,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	jsonData, err := actorResponses.MarshalJSON()
	if err != nil {
		http.Error(w, "response error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func (h *MovieHandler) GetUserRating(w http.ResponseWriter, r *http.Request) {
	logger := log.Ctx(r.Context())

	movieIDStr := r.URL.Query().Get("movie_id")

	movieID, err := strconv.Atoi(movieIDStr)
	if err != nil {
		api.Response(r.Context(), w, http.StatusBadRequest, api.PreparedDefaultError("bad_request", err))
		return
	}

	rating, errServResp := h.movieService.GetUserRating(r.Context(), movieID)
	if errServResp != nil {
		logger.Error().Err(errServResp.Error).Msg("failed to get user rating")
		api.Response(r.Context(), w, http.StatusInternalServerError, api.PreparedDefaultError("internal_error", err))
		return
	}

	api.Response(r.Context(), w, http.StatusOK, map[string]int{"rating": rating})
}

func (h *MovieHandler) AddOrUpdateRating(w http.ResponseWriter, r *http.Request) {
	logger := log.Ctx(r.Context())

	mvID, err := strconv.Atoi(mux.Vars(r)["movie_id"])
	if err != nil {
		errMsg := fmt.Errorf("getMovie action: Bad request - %w", err)
		logger.Error().Err(errMsg).Msg("bad_request")
		api.Response(r.Context(), w, http.StatusBadRequest, api.PreparedDefaultError("bad_request", errMsg))

		return
	}

	req := &api.AddOrUpdateRatingReq{}
	if !api.DecodeBody(w, r, req) {
		return
	}

	if req.Rating < 1 || req.Rating > 10 {
		api.Response(r.Context(), w, http.StatusBadRequest, api.PreparedDefaultError("bad_request", errors.New("rating must be between 1 and 10")))
		return
	}

	if errServResp := h.movieService.AddOrUpdateRating(r.Context(), mvID, req.Rating); errServResp != nil {
		logger.Error().Err(errServResp.Error).Msg("failed to add or update rating")
		api.Response(r.Context(), w, http.StatusInternalServerError, api.PreparedDefaultError("internal_error", err))
		return
	}

	api.Response(r.Context(), w, http.StatusOK, map[string]string{"message": "rating updated"})
}

func (h *MovieHandler) DeleteRating(w http.ResponseWriter, r *http.Request) {
	logger := log.Ctx(r.Context())

	mvID, err := strconv.Atoi(mux.Vars(r)["movie_id"])
	if err != nil {
		errMsg := fmt.Errorf("getMovie action: Bad request - %w", err)
		logger.Error().Err(errMsg).Msg("bad_request")
		api.Response(r.Context(), w, http.StatusBadRequest, api.PreparedDefaultError("bad_request", errMsg))

		return
	}

	errServResp := h.movieService.DeleteRating(r.Context(), mvID)
	if errServResp != nil {
		log.Print("errServResp", errServResp)
		http.Error(w, "failed to delete rating", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
