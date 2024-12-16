package service

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

// SearchActors поиск актеров в эластике
func (s *MovieService) SearchActors(ctx context.Context, query string) ([]models.ActorInfo, error) {
	return s.movieClient.SearchActors(ctx, query)
}
