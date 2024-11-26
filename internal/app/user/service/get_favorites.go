package service

import (
	"context"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
)

func (u *UserService) GetFavorites(ctx context.Context, usrID int) ([]models.MovieShortInfo, *errVals.ServiceError) {
	// TODO: Uncomment when movie_service microservice will be ready
	// mvIDs, err := u.userClient.GetFavorites(ctx, usrID)
	// if err != nil {
	// 	return nil, errVals.NewServiceError("failed_to_get_user_favorites", err)
	// }

	// request to movie_service microservice

	// return repResp, nil
	return nil, nil
}
