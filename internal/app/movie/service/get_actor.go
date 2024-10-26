package service

import (
	"context"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

func (s *MovieService) GetActor(ctx context.Context, actorId int) (*models.ActorInfo, *models.ErrorRespData) {
	actor, err, code := s.movieRepository.GetActor(ctx, actorId)

	if err != nil {
		errs := make([]errVals.ErrorObj, 1)
		errs[0] = *err

		return nil, &models.ErrorRespData{
			StatusCode: code,
			Errors:     errs,
		}
	}

	return actor, nil
}
