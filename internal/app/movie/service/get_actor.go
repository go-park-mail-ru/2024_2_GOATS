package service

import (
	"context"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

func (s *MovieService) GetActor(ctx context.Context, actorID int) (*models.ActorInfo, *errVals.ServiceError) {
	// actor, err := s.movieRepository.GetActor(ctx, actorID)

	// if err != nil {
	// 	return nil, errVals.ToServiceErrorFromRepo(err)
	// }

	// return actor, nil
	return nil, nil
}
