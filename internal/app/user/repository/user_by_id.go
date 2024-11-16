package repository

import (
	"context"
	"database/sql"
	"errors"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/repository/converter"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/repository/userdb"
)

func (u *UserRepo) UserByID(ctx context.Context, userID int) (*models.User, *errVals.RepoError) {
	usr, err := userdb.FindByID(ctx, userID, u.Database)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errVals.NewRepoError(errVals.ErrUserNotFoundCode, errVals.ErrUserNotFound)
		}

		return nil, errVals.NewRepoError(errVals.ErrServerCode, errVals.NewCustomError(err.Error()))
	}

	return converter.ToUserFromDBUser(usr), nil
}
