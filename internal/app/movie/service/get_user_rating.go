package service

import (
	"context"
	"errors"
	"math"

	"github.com/go-park-mail-ru/2024_2_GOATS/config"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
)

// GetUserRating получение рейтинга
func (s *MovieService) GetUserRating(ctx context.Context, movieID int32) (int32, *errVals.ServiceError) {

	usrID := config.CurrentUserID(ctx)
	if usrID < 0 {
		return 0, &errVals.ServiceError{
			Code:  "USER_ZERO",
			Error: errors.New("error usrID = 0"),
		}
	}

	if usrID < math.MinInt32 || usrID > math.MaxInt32 {
		return 0, &errVals.ServiceError{
			Code:  "OUT_OF_INTERVAL",
			Error: errors.New("invalid usrID"),
		}
	}

	rating, err := s.movieClient.GetUserRating(ctx, movieID, int32(usrID))

	if err != nil {
		return 0, nil
	}

	return rating, nil
}
