package delivery

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/handlers"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

var _ handlers.MovieImplementationInterface = (*Implementation)(nil)

//go:generate mockgen -source=implementation.go -destination=mocks/mock.go
type MovieServiceInterface interface {
	GetCollection(ctx context.Context) (*models.CollectionsResponse, *models.ErrorResponse)
}

type Implementation struct {
	Ctx          context.Context
	movieService MovieServiceInterface
}

func NewImplementation(ctx context.Context, srv MovieServiceInterface) *Implementation {
	return &Implementation{
		Ctx:          ctx,
		movieService: srv,
	}
}
