package delivery

import (
	"context"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api/converter"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api/handlers"
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
