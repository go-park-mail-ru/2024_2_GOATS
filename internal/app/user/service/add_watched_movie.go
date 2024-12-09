package service

import (
	"context"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

func (u *UserService) AddWatchedMovie(ctx context.Context, watchedData *models.OwnWatchedMovie) *errVals.ServiceError {
	err := u.userClient.AddWatchedMovie(ctx, watchedData)
	if err != nil {
		return errVals.NewServiceError("failed_to_set_favorite", err)
	}

	return nil
}
