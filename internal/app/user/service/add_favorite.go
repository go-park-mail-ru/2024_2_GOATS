package service

import (
	"context"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/service/converter"
)

func (u *UserService) AddFavorite(ctx context.Context, favData *models.Favorite) *errVals.ServiceError {
	err := u.userRepo.CreateFavorite(ctx, converter.ToRepoFavoriteFromFavorite(favData))
	if err != nil {
		return errVals.ToServiceErrorFromRepo(err)
	}

	return nil
}
