package delivery

import (
	"net/http"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api/converter"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api/handlers"
)

var _ handlers.MovieImplementationInterface = (*MovieHandler)(nil)

type MovieHandler struct {
	movieService MovieServiceInterface
	cfg          *config.Config
}

func NewMovieHandler(srv MovieServiceInterface, cfg *config.Config) *MovieHandler {
	return &MovieHandler{
		movieService: srv,
		cfg:          cfg,
	}
}

func (m *MovieHandler) GetCollections(w http.ResponseWriter, r *http.Request) {
	ctx := config.WrapContext(r.Context(), m.cfg)
	collectionsServResp, errServResp := m.movieService.GetCollection(ctx)
	collectionsResp, errResp := converter.ToApiCollectionsResponse(collectionsServResp), converter.ToApiErrorResponse(errServResp)

	if errResp != nil {
		api.Response(w, errResp.StatusCode, errResp)
		return
	}

	api.Response(w, collectionsResp.StatusCode, collectionsResp)
}
