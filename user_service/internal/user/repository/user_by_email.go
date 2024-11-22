package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/repository/converter"
	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/repository/userdb"
	srvDTO "github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/service/dto"
)

func (u *UserRepo) UserByEmail(ctx context.Context, email string) (*srvDTO.User, error) {
	usr, err := userdb.FindByEmail(ctx, email, u.Database)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New(fmt.Sprintf(errVals.ErrUserNotFoundCode, err))
		}

		return nil, fmt.Errorf("%s: %w", errVals.ErrServerCode, err)
	}

	return converter.ToUserFromRepoUser(usr), nil
}
