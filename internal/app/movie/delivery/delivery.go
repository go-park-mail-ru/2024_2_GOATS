package delivery

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

//go:generate mockgen -source=delivery.go -destination=mocks/mock.go
type MovieServiceInterface interface {
	GetCollection(ctx context.Context) (*models.CollectionsRespData, *models.ErrorRespData)
	GetMovie(ctx context.Context, mvId int) (*models.MovieFullData, *models.ErrorRespData)
}
