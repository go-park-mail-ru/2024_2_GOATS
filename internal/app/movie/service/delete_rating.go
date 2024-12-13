package service

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/config"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"log"
)

func (s *MovieService) DeleteRating(ctx context.Context, movieID int) *errVals.ServiceError {

	usrID := config.CurrentUserID(ctx)
	if usrID == 0 {
		return &errVals.ServiceError{
			Code:  "USER_ZERO",
			Error: errors.New("error usrID = 0"),
		}
	}

	log.Println("ServiceDEL", movieID, usrID)

	err := s.movieClient.DeleteUserRating(ctx, movieID, usrID)

	if err != nil {
		return &errVals.ServiceError{
			Code:  "RATING_DELETE_ERROR",
			Error: errors.New("internal server error"),
		}
	}

	return nil
}
