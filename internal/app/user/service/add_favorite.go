package service

import (
	"context"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

// AddFavorite sets new user favorite by calling userClient SetFavorite
func (u *UserService) AddFavorite(ctx context.Context, favData *models.Favorite) *errVals.ServiceError {
	err := u.userClient.SetFavorite(ctx, favData)
	if err != nil {
		return errVals.NewServiceError("failed_to_set_favorite", err)
	}

	return nil
}
