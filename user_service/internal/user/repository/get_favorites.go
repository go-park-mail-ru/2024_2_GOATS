package repository

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/repository/favoritedb"
)

func (r *UserRepo) GetFavorites(ctx context.Context, usrID uint64) ([]uint64, error) {
	rows, err := favoritedb.FindByUserID(ctx, usrID, r.Database)
	if err != nil {
		// return nil, fmt.Errorf(err)
	}

	favorites, err := favoritedb.ScanConnections(rows)
	if err != nil {
		// return nil, errVals.NewRepoError(errVals.ErrServerCode, errVals.NewCustomError(err.Error()))
	}

	return favorites, nil
}
