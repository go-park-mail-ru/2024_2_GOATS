package repository

import (
	"context"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/repository/converter"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/repository/favoritedb"
)

func (r *UserRepo) GetFavorites(ctx context.Context, usrID int) ([]models.MovieShortInfo, *errVals.RepoError) {
	rows, err := favoritedb.FindByUserID(ctx, usrID, r.Database)
	if err != nil {
		return nil, errVals.NewRepoError(errVals.ErrGetFavorites, errVals.NewCustomError(err.Error()))
	}

	favorites, err := favoritedb.ScanConnections(rows)
	if err != nil {
		return nil, errVals.NewRepoError(errVals.ErrServerCode, errVals.NewCustomError(err.Error()))
	}

	var favs []models.MovieShortInfo

	for _, fav := range favorites {
		favs = append(favs, *converter.ToMovieShortInfoFromRepo(fav))
	}

	return favs, nil
}
