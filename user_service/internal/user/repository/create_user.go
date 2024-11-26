package repository

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/repository/converter"
	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/repository/dto"
	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/repository/password"
	"github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/repository/userdb"
	srvDTO "github.com/go-park-mail-ru/2024_2_GOATS/user_service/internal/user/service/dto"
)

func (u *UserRepo) CreateUser(ctx context.Context, registerData *dto.RepoCreateData) (*srvDTO.User, error) {
	hashedPasswd, err := password.HashAndSalt(ctx, registerData.Password)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errors.ErrCreateUserCode, err)
	}

	registerData.Password = hashedPasswd

	usr, err := userdb.Create(ctx, *registerData, u.Database)
	if err != nil {
		if errors.IsDuplicateError(err) {
			return nil, fmt.Errorf("%s: duplicate entry: %w", errors.DuplicateErrCode, err)
		}

		return nil, fmt.Errorf("%s: %w", errors.ErrCreateUserCode, err)
	}

	return converter.ToUserFromRepoUser(usr), nil
}
