package service

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/config"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
)

func (s *MovieService) AddOrUpdateRating(ctx context.Context, movieID, rating int) *errVals.ServiceError {
	usrID := config.CurrentUserID(ctx)
	if usrID == 0 {
		return &errVals.ServiceError{
			Code:  "USER_ZERO",
			Error: errors.New("error usrID = 0"),
		}
	}

	err := s.movieClient.AddOrUpdateRating(ctx, movieID, usrID, rating)
	if err != nil {
		return &errVals.ServiceError{
			Code:  "ADD_RATING_ERROR",
			Error: errors.New("internal server error"),
		}
	}

	return nil
}
