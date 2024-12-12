package service

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/config"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
)

func (s *MovieService) GetUserRating(ctx context.Context, movieID int) (int, *errVals.ServiceError) {

	usrID := config.CurrentUserID(ctx)
	if usrID == 0 {
		return 0, &errVals.ServiceError{
			Code:  "USER_ZERO",
			Error: errors.New("error usrID = 0"),
		}
	}

	rating, err := s.movieClient.GetUserRating(ctx, movieID, usrID)

	if err != nil {
		return 0, nil
	}

	return rating, nil
}
