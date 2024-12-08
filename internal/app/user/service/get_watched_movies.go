package service

import (
	"context"
	"fmt"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

func (u *UserService) GetWatchedMovies(ctx context.Context, usrID int) ([]models.WatchedMovieInfo, *errVals.ServiceError) {
	movies, err := u.userClient.GetWatchedMovies(ctx, usrID)
	if err != nil {
		err = fmt.Errorf("fail in movie_service %w", err)
		return nil, errVals.NewServiceError("failed_to_get_user_favorites", err)
	}

	return movies, nil
}
