package service

import (
	"context"
	"fmt"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

func (s *MovieService) GetCollection(ctx context.Context, filter string) (*models.CollectionsRespData, *errVals.ServiceError) {
	collections, err := s.movieClient.GetCollection(ctx, filter)

	if err != nil {
		return nil, &errVals.ServiceError{
			Code:  "GetCollection",
			Error: fmt.Errorf("error GetCollection: %w", err),
		}
	}

	return &models.CollectionsRespData{Collections: collections}, nil
}
