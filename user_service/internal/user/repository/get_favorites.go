package repository

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/repository/favoritedb"
)

func (r *UserRepo) GetFavorites(ctx context.Context, usrID uint64) ([]uint64, error) {
	rows, err := favoritedb.FindByUserID(ctx, usrID, r.Database)
	if err != nil {
		return nil, fmt.Errorf("cannot_get_favorites: %w", err)
	}

	favorites, err := favoritedb.ScanConnections(rows)
	if err != nil {
		return nil, fmt.Errorf("cannot_scan_favorites: %w", err)
	}

	return favorites, nil
}
