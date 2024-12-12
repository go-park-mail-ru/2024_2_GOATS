package service

import (
	"context"
	"errors"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

// GetActor gets actor by id
func (s *MovieService) GetActor(ctx context.Context, actorID int) (*models.ActorInfo, *errVals.ServiceError) {
	actor, err := s.movieClient.GetActor(ctx, actorID)

	if err != nil {
		return nil, &errVals.ServiceError{
			Code:  "ACTOR_NOT_FOUND",
			Error: errors.New("internal server error"),
		}
	}

	return actor, nil
}
