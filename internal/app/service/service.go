package service

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/api"
	errVals "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/errors"
	"github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models"
	authModels "github.com/go-park-mail-ru/2024_2_GOATS/internal/app/models/auth"
)

var _ api.ServiceInterface = (*Service)(nil)

type RepositoryInterface interface {
	Login(ctx context.Context, loginData *authModels.LoginData) ([]*authModels.CookieData, *errVals.ErrorObj, int)
	Register(ctx context.Context, registerData *authModels.RegisterData) (*authModels.CookieData, *errVals.ErrorObj, int)
	GetCollection(ctx context.Context) ([]models.Collection, *errVals.ErrorObj, int)
	Session(ctx context.Context, cookie string) (*models.User, *errVals.ErrorObj, int)
	Logout(ctx context.Context, cookie string) (*authModels.CookieData, *errVals.ErrorObj, int)
}

type Service struct {
	repository RepositoryInterface
}

func NewService(repo RepositoryInterface) api.ServiceInterface {
	return &Service{
		repository: repo,
	}
}
