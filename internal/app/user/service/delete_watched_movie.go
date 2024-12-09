package service

import (
	"context"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

func (u *UserService) DeleteWatchedMovie(ctx context.Context, watchedData *models.DeletedWatchedMovie) *errVals.ServiceError {
	err := u.userClient.DeleteWatchedMovie(ctx, watchedData)
	if err != nil {
		return errVals.NewServiceError("failed_to_set_favorite", err)
	}

	return nil
}
