package service

import (
	"context"
	"fmt"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

func (u *UserService) GetFavorites(ctx context.Context, usrID int) ([]models.MovieShortInfo, *errVals.ServiceError) {
	mvIDs, err := u.userClient.GetFavorites(ctx, usrID)
	if err != nil {
		return nil, errVals.NewServiceError("failed_to_get_user_favorites", err)
	}

	movies, err := u.mvClient.GetFavorites(ctx, mvIDs)
	if err != nil {
		err = fmt.Errorf("fail in movie_service %w", err)
		return nil, errVals.NewServiceError("failed_to_get_user_favorites", err)
	}

	return movies, nil
}
