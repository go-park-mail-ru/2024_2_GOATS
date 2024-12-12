package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/repository/converter"
	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/repository/subscriptiondb"
	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/repository/userdb"
	srvDTO "github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/service/dto"
)

// UserByID finds user by id by calling db FindByID
func (u *UserRepo) UserByID(ctx context.Context, userID uint64) (*srvDTO.User, error) {
	usr, err := userdb.FindByID(ctx, userID, u.Database)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New(fmt.Sprint(errVals.ErrUserNotFoundCode, err))
		}

		return nil, fmt.Errorf("%s: %w", errVals.ErrServerCode, err)
	}

	subs, err := subscriptiondb.FindByUserID(ctx, usr.ID, u.Database)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errVals.ErrServerCode, err)
	}

	return converter.ToUserFromRepoUser(usr, subs), nil
}
