package service

import (
	"context"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

// ResetFavorite reset user favorite by calling userClient ResetFavorite
func (u *UserService) ResetFavorite(ctx context.Context, favData *models.Favorite) *errVals.ServiceError {
	err := u.userClient.ResetFavorite(ctx, favData)
	if err != nil {
		return errVals.NewServiceError("failed_to_reset_favorite", err)
	}

	return nil
}
