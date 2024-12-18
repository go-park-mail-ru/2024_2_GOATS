package service

import (
	"context"
	"errors"
	"math"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
)

// AddOrUpdateRating creates user_rating for movie
func (s *MovieService) AddOrUpdateRating(ctx context.Context, movieID, rating int32) *errVals.ServiceError {
	usrID := config.CurrentUserID(ctx)
	if usrID == 0 {
		return &errVals.ServiceError{
			Code:  "USER_ZERO",
			Error: errors.New("error usrID = 0"),
		}
	}

	if usrID < math.MinInt32 || usrID > math.MaxInt32 {
		return &errVals.ServiceError{
			Code:  "OUT_OF_INTERVAL",
			Error: errors.New("invalid usrID"),
		}
	}

	err := s.movieClient.AddOrUpdateRating(ctx, movieID, int32(usrID), rating)
	if err != nil {
		return &errVals.ServiceError{
			Code:  "ADD_RATING_ERROR",
			Error: errors.New("internal server error"),
		}
	}

	return nil
}
