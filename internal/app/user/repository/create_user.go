package repository

import (
	"context"
	"fmt"

	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/repository/converter"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/repository/dto"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/repository/password"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/user/repository/userdb"
)

func (u *UserRepo) CreateUser(ctx context.Context, registerData *dto.DBRegisterData) (*models.User, *errVals.RepoError) {
	hashedPasswd, err := password.HashAndSalt(ctx, registerData.Password)
	if err != nil {
		return nil, errVals.NewRepoError(
			errVals.ErrServerCode,
			errVals.NewCustomError(fmt.Sprintf("error hashing password: %v", err)),
		)
	}

	registerData.Password = hashedPasswd

	usr, err := userdb.Create(ctx, *registerData, u.Database)
	if err != nil {
		return nil, errVals.NewRepoError(errVals.ErrCreateUserCode, errVals.NewCustomError(err.Error()))
	}

	return converter.ToUserFromDBUser(usr), nil
}
