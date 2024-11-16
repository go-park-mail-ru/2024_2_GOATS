package service

import (
	"context"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

func (u *UserService) GetFavorites(ctx context.Context, usrID int) ([]models.MovieShortInfo, *errVals.ServiceError) {
	repResp, err := u.userRepo.GetFavorites(ctx, usrID)
	if err != nil {
		return nil, errVals.ToServiceErrorFromRepo(err)
	}

	return repResp, nil
}
