package service

import (
	"context"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

func (s *Service) GetCollection(ctx context.Context) (*models.CollectionsResponse, *models.ErrorResponse) {
	collections, err, code := s.repository.GetCollection(ctx)

	if err != nil {
		errors := make([]errVals.ErrorObj, 1)
		errors[0] = *err

		return nil, &models.ErrorResponse{
			Success:    false,
			StatusCode: code,
			Errors:     errors,
		}
	}

	return &models.CollectionsResponse{
		Collections: collections,
		Success:     true,
	}, nil
}
